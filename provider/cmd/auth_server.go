package cmd

import (
	"bytes"
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"io"
	"net/http"
	"time"

	"golang.org/x/sync/errgroup"

	sdkclient "github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/ovrclk/akash/cmd/common"
	gwrest "github.com/ovrclk/akash/provider/gateway/rest"
	cmodule "github.com/ovrclk/akash/x/cert"
	cutils "github.com/ovrclk/akash/x/cert/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	FlagJwtAuthListenAddress = "jwt-auth-listen-address"
	FlagJwtExpiresAfter      = "jwt-expires-after"
)

func AuthServerCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:          "auth-server",
		Short:        "Run the authentication server to issue JWTs to tenants",
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			return common.RunForeverWithContext(cmd.Context(), func(ctx context.Context) error {
				return doAuthServerCmd(ctx, cmd, args)
			})
		},
	}
	flags.AddTxFlagsToCmd(cmd)

	cmd.Flags().String(FlagJwtAuthListenAddress, "0.0.0.0:8444", "`host:port` for the JWT server to listen on")
	if err := viper.BindPFlag(FlagJwtAuthListenAddress, cmd.Flags().Lookup(FlagJwtAuthListenAddress)); err != nil {
		return nil
	}

	cmd.Flags().Duration(FlagJwtExpiresAfter, 600*time.Second, "duration for which the JWT is valid after it is issued")
	if err := viper.BindPFlag(FlagJwtExpiresAfter, cmd.Flags().Lookup(FlagJwtExpiresAfter)); err != nil {
		return nil
	}

	cmd.Flags().String(FlagAuthPem, "", "")

	return cmd
}

func doAuthServerCmd(ctx context.Context, cmd *cobra.Command, _ []string) error {
	expiresAfter := viper.GetDuration(FlagJwtExpiresAfter)
	jwtGwAddr := viper.GetString(FlagJwtAuthListenAddress)

	cctx, err := sdkclient.GetClientTxContext(cmd)
	if err != nil {
		return err
	}

	txFactory := tx.NewFactoryCLI(cctx, cmd.Flags()).WithTxConfig(cctx.TxConfig).WithAccountRetriever(cctx.AccountRetriever)

	var certFromFlag io.Reader
	if val := cmd.Flag(FlagAuthPem).Value.String(); val != "" {
		certFromFlag = bytes.NewBufferString(val)
	}

	cpem, err := cutils.LoadPEMForAccount(cctx, txFactory.Keybase(), cutils.PEMFromReader(certFromFlag))
	if err != nil {
		return err
	}

	blk, _ := pem.Decode(cpem.Cert)
	x509cert, err := x509.ParseCertificate(blk.Bytes)
	if err != nil {
		return err
	}

	cert, err := tls.X509KeyPair(cpem.Cert, cpem.Priv)
	if err != nil {
		return err
	}

	cquery := cmodule.AppModuleBasic{}.GetQueryClient(cctx)
	group, ctx := errgroup.WithContext(ctx)

	jwtGateway, err := gwrest.NewJwtServer(
		ctx,
		cquery,
		jwtGwAddr,
		cctx.FromAddress,
		cert,
		x509cert.SerialNumber.String(),
		expiresAfter,
	)
	if err != nil {
		return err
	}

	group.Go(func() error {
		return jwtGateway.ListenAndServeTLS("", "")
	})

	group.Go(func() error {
		<-ctx.Done()
		return jwtGateway.Close()
	})

	err = group.Wait()
	if err != nil && !errors.Is(err, context.Canceled) && !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	return nil
}
