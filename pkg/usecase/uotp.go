package usecase

import (
	"context"
	"fmt"

	"github.com/Akshayvij07/ecommerce/pkg/config"
	"github.com/Akshayvij07/ecommerce/pkg/helper/request"
	interfaces "github.com/Akshayvij07/ecommerce/pkg/usecase/interface"
	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/verify/v2"
)

type OtpUseCase struct {
	cfg config.Config
}

func NewOtpUseCase(cfg config.Config) interfaces.OtpUseCase {
	return &OtpUseCase{

		cfg: cfg,
	}
}

func (c *OtpUseCase) SendOTP(ctx context.Context, mobno request.OTPreq) (string, error) {

	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Password: c.cfg.AUTHTOCKEN,
		Username: c.cfg.ACCOUNTSID,
	})

	fmt.Println(mobno, c.cfg.AUTHTOCKEN, c.cfg.ACCOUNTSID, c.cfg.SERVICES_ID)
	params := &twilioApi.CreateVerificationParams{}
	params.SetTo(mobno.Phone)
	params.SetChannel("sms")

	resp, err := client.VerifyV2.CreateVerification(c.cfg.SERVICES_ID, params)
	if err != nil {
		return "", err
	}

	return *resp.Sid, nil
}

func (c *OtpUseCase) VerifyOTP(ctx context.Context, userData request.Otpverifier) error {

	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Password: c.cfg.AUTHTOCKEN,
		Username: c.cfg.ACCOUNTSID,
	})

	params := &twilioApi.CreateVerificationCheckParams{}
	params.SetTo(userData.Phone)
	params.SetCode(userData.Pin)
	resp, err := client.VerifyV2.CreateVerificationCheck(c.cfg.SERVICES_ID, params)
	if err != nil {
		return err
	} else if *resp.Status == "approved" {
		return nil
	}

	return nil
}
