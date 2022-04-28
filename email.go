package main

import (
	"bytes"
	"strconv"
	"text/template"
	"time"

	gomail "gopkg.in/gomail.v2"
)

const rawEmailTemplate string = `<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional //EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
    <html
      xmlns="http://www.w3.org/1999/xhtml"
      xmlns:o="urn:schemas-microsoft-com:office:office"
      xmlns:v="urn:schemas-microsoft-com:vml"
    >
      <head>
        <!--[if gte mso 9
          ]><xml
            ><o:OfficeDocumentSettings
              ><o:AllowPNG /><o:PixelsPerInch
                >96</o:PixelsPerInch
              ></o:OfficeDocumentSettings
            ></xml
          ><!
        [endif]-->
        <meta content="text/html; charset=utf-8" http-equiv="Content-Type" />
        <meta content="width=device-width" name="viewport" />
        <!--[if !mso]><!-->
        <meta content="IE=edge" http-equiv="X-UA-Compatible" />
        <!--<![endif]-->
        <title></title>
        <!--[if !mso]><!-->
        <link
          href="https://fonts.googleapis.com/css?family=Roboto"
          rel="stylesheet"
          type="text/css"
        />
        <!--<![endif]-->

        <style type="text/css">
          @import url('https://fonts.googleapis.com/css2?family=Arvo&family=Merriweather:wght@400;700&display=swap');

          body {
            margin: 0;
            padding: 0;
          }
    
          table,
          td,
          tr {
            vertical-align: top;
            border-collapse: collapse;
          }
    
          * {
            line-height: inherit;
          }
    
          a[x-apple-data-detectors="true"] {
            color: inherit !important;
            text-decoration: none !important;
          }

          .postTitle{
            font-family: "Merriweather", "Bookerly",Georgia,serif; font-size: 16px;
          }

          .postSource{
            color: gray;
          }

          .postDescription{
            font-family: "Merriweather", "Bookerly",Georgia,serif; font-size: 16px; margin-bottom: 48px;
          }

          .collectionTitle{
            text-align: center;
            font-size: 26px;
            font-weight: 400;
            font-family: "Arlo", serif;
          }

          .newsletterTitle{
            text-align: left;
          }

          .newsletterDate{
            text-align: right;
            color: #888;
          }

          .headerBar{
            font-family: "Arlo", serif;
            margin-top: 48px;
            width: 100%; table-layout: fixed; vertical-align: top;

            font-size: 24px;
            font-family: monospace;

            max-width: 800px;
            margin: 0 auto;
          }

          a{
            color: black !important;
            text-decoration: none !important;
          }
        </style>
        <style id="media-query" type="text/css">
          @media (max-width: 670px) {
            .block-grid,
            .col {
              min-width: 320px !important;
              max-width: 100% !important;
              display: block !important;
            }
    
            .block-grid {
              width: 100% !important;
            }
    
            .col {
              width: 100% !important;
            }
    
            .col > div {
              margin: 0 auto;
            }
    
            img.fullwidth,
            img.fullwidthOnMobile {
              max-width: 100% !important;
            }
    
            .no-stack .col {
              min-width: 0 !important;
              display: table-cell !important;
            }
    
            .no-stack.two-up .col {
              width: 50% !important;
            }
    
            .no-stack .col.num4 {
              width: 33% !important;
            }
    
            .no-stack .col.num8 {
              width: 66% !important;
            }
    
            .no-stack .col.num4 {
              width: 33% !important;
            }
    
            .no-stack .col.num3 {
              width: 25% !important;
            }
    
            .no-stack .col.num6 {
              width: 50% !important;
            }
    
            .no-stack .col.num9 {
              width: 75% !important;
            }
    
            .video-block {
              max-width: none !important;
            }
    
            .mobile_hide {
              min-height: 0px;
              max-height: 0px;
              max-width: 0px;
              display: none;
              overflow: hidden;
              font-size: 0px;
            }
    
            .desktop_hide {
              display: block !important;
              max-height: none !important;
            }
          }
        </style>
      </head>
      <body
        class="clean-body"
        style="
          margin: 0;
          padding: 0;
          -webkit-text-size-adjust: 100%;
          background-color: #ffffff;
        "
      >
        <!--[if IE]><div class="ie-browser"><![endif]-->
        <table
          bgcolor="#FFFFFF"
          cellpadding="0"
          cellspacing="0"
          class="nl-container"
          role="presentation"
          style="
            table-layout: fixed;
            vertical-align: top;
            min-width: 320px;
            margin: 0 auto;
            border-spacing: 0;
            border-collapse: collapse;
            mso-table-lspace: 0pt;
            mso-table-rspace: 0pt;
            background-color: #ffffff;
            width: 100%;
          "
          valign="top"
          width="100%"
        >
          <tbody>
            <tr>
              <td>
                <table class="headerBar">
                  <tr>
                    <td class="newsletterTitle">The Weekly Digest</td>
                    <td class="newsletterDate">{{.Date}}</td>
                  </tr>
                </table>
              </td>
            </tr>

            <tr style="vertical-align: top;" valign="top">
              <td style="word-break: break-word; vertical-align: top;" valign="top">
                <!--[if (mso)|(IE)]><table width="100%" cellpadding="0" cellspacing="0" border="0"><tr><td align="center" style="background-color:#FFFFFF"><![endif]-->
                <div style="background-color: #fff;">
                  <div
                    class="block-grid"
                    style="
                      margin: 0 auto;
                      min-width: 320px;
                      max-width: 650px;
                      overflow-wrap: break-word;
                      word-wrap: break-word;
                      word-break: break-word;
                      background-color: transparent;
                    "
                  >
                    <div
                      style="
                        border-collapse: collapse;
                        display: table;
                        width: 100%;
                        background-color: transparent;
                      "
                    >
                      <!--[if (mso)|(IE)]><table width="100%" cellpadding="0" cellspacing="0" border="0" style="background-color:#fff;"><tr><td align="center"><table cellpadding="0" cellspacing="0" border="0" style="width:650px"><tr class="layout-full-width" style="background-color:transparent"><![endif]-->
                      <!--[if (mso)|(IE)]><td align="center" width="650" style="background-color:transparent;width:650px; border-top: 0px solid transparent; border-left: 0px solid transparent; border-bottom: 0px solid transparent; border-right: 0px solid transparent;" valign="top"><table width="100%" cellpadding="0" cellspacing="0" border="0"><tr><td style="padding-right: 0px; padding-left: 0px; padding-top:25px; padding-bottom:25px;"><![endif]-->
                      <div
                        class="col num12"
                        style="
                          min-width: 320px;
                          max-width: 650px;
                          display: table-cell;
                          vertical-align: top;
                          width: 650px;
                        "
                      >
                        <div style="width: 100% !important;">
                          <!--[if (!mso)&(!IE)]><!-->
                          <div
                            style="
                              border-top: 0px solid transparent;
                              border-left: 0px solid transparent;
                              border-bottom: 0px solid transparent;
                              border-right: 0px solid transparent;
                              padding-top: 25px;
                              padding-bottom: 25px;
                              padding-right: 0px;
                              padding-left: 0px;
                            "
                          >
                            <!--<![endif]-->
                            <!--[if mso]><table width="100%" cellpadding="0" cellspacing="0" border="0"><tr><td style="padding-right: 25px; padding-left: 25px; padding-top: 25px; padding-bottom: 25px; font-family: Tahoma, Verdana, sans-serif"><![endif]-->
                            <div
                              style="
                                color: #000000;
                                font-family: 'Roboto', Tahoma, Verdana, Segoe,
                                  sans-serif;
                                line-height: 1.5;
                                padding-top: 25px;
                                padding-right: 25px;
                                padding-bottom: 25px;
                                padding-left: 25px;
                              "
                            >
                              <div
                                style="
                                  line-height: 1.5;
                                  font-size: 12px;
                                  font-family: 'Roboto', Tahoma, Verdana, Segoe,
                                    sans-serif;
                                  color: #000000;
                                  mso-line-height-alt: 18px;
                                "
                              >
                                <p
                                  style="
                                    font-size: 16px;
                                    line-height: 1.5;
                                    word-break: break-word;
                                    text-align: left;
                                    font-family: 'Roboto', Tahoma, Verdana, Segoe,
                                      sans-serif;
                                    mso-line-height-alt: 24px;
                                    margin: 0;
                                  "
                                >
                                  <span style="font-size: 16px;"
                                    >{{.HtmlContent}}</span
                                  >
                                </p>
                              </div>
                            </div>
                            <!--[if mso]></td></tr></table><![endif]-->
                            <!--[if (!mso)&(!IE)]><!-->
                          </div>
                          <!--<![endif]-->
                        </div>
                      </div>
                      <!--[if (mso)|(IE)]></td></tr></table><![endif]-->
                      <!--[if (mso)|(IE)]></td></tr></table></td></tr></table><![endif]-->
                    </div>
                  </div>
                </div>
                <div style="background-color: #f2f1f1;">
                  <div
                    class="block-grid"
                    style="
                      margin: 0 auto;
                      min-width: 320px;
                      max-width: 650px;
                      overflow-wrap: break-word;
                      word-wrap: break-word;
                      word-break: break-word;
                      background-color: transparent;
                    "
                  >
                    <div
                      style="
                        border-collapse: collapse;
                        display: table;
                        width: 100%;
                        background-color: transparent;
                      "
                    >
                      <!--[if (mso)|(IE)]><table width="100%" cellpadding="0" cellspacing="0" border="0" style="background-color:#f2f1f1;"><tr><td align="center"><table cellpadding="0" cellspacing="0" border="0" style="width:650px"><tr class="layout-full-width" style="background-color:transparent"><![endif]-->
                      <!--[if (mso)|(IE)]><td align="center" width="650" style="background-color:transparent;width:650px; border-top: 0px solid transparent; border-left: 0px solid transparent; border-bottom: 0px solid transparent; border-right: 0px solid transparent;" valign="top"><table width="100%" cellpadding="0" cellspacing="0" border="0"><tr><td style="padding-right: 0px; padding-left: 0px; padding-top:30px; padding-bottom:30px;"><![endif]-->
                      <div
                        class="col num12"
                        style="
                          min-width: 320px;
                          max-width: 650px;
                          display: table-cell;
                          vertical-align: top;
                          width: 650px;
                        "
                      >
                        <div style="width: 100% !important;">
                          <!--[if (!mso)&(!IE)]><!-->
                          <div
                            style="
                              border-top: 0px solid transparent;
                              border-left: 0px solid transparent;
                              border-bottom: 0px solid transparent;
                              border-right: 0px solid transparent;
                              padding-top: 30px;
                              padding-bottom: 30px;
                              padding-right: 0px;
                              padding-left: 0px;
                            "
                          >
                            <!--<![endif]-->
                            <!--[if mso]><table width="100%" cellpadding="0" cellspacing="0" border="0"><tr><td style="padding-right: 25px; padding-left: 25px; padding-top: 25px; padding-bottom: 25px; font-family: Tahoma, Verdana, sans-serif"><![endif]-->
                            <div
                              style="
                                color: #9a9999;
                                font-family: 'Roboto', Tahoma, Verdana, Segoe,
                                  sans-serif;
                                line-height: 1.5;
                                padding-top: 25px;
                                padding-right: 25px;
                                padding-bottom: 25px;
                                padding-left: 25px;
                              "
                            >
                              <div
                                style="
                                  line-height: 1.5;
                                  font-size: 12px;
                                  font-family: 'Roboto', Tahoma, Verdana, Segoe,
                                    sans-serif;
                                  color: #9a9999;
                                  mso-line-height-alt: 18px;
                                "
                              >
                                <p
                                  style="
                                    font-size: 12px;
                                    line-height: 1.5;
                                    word-break: break-word;
                                    text-align: center;
                                    font-family: 'Roboto', Tahoma, Verdana, Segoe,
                                      sans-serif;
                                    mso-line-height-alt: 18px;
                                    margin: 0;
                                  "
                                >
                                  <span style="font-size: 12px;"
                                    >You are receiving this email because you signed up for the Feeds to Email service
                                    <br/>
                                    <br/>
                                  </span>
                                </p>
                                <p
                                  style="
                                    font-size: 12px;
                                    line-height: 1.5;
                                    word-break: break-word;
                                    text-align: center;
                                    font-family: 'Roboto', Tahoma, Verdana, Segoe,
                                      sans-serif;
                                    mso-line-height-alt: 18px;
                                    margin: 0;
                                  "
                                >
                                  <span style="font-size: 12px;"
                                    >{{.Company}}, {{.Address}}</span
                                  >
                                </p>
                                <p
                                  style="
                                    font-size: 12px;
                                    line-height: 1.5;
                                    word-break: break-word;
                                    text-align: center;
                                    font-family: 'Roboto', Tahoma, Verdana, Segoe,
                                      sans-serif;
                                    mso-line-height-alt: 18px;
                                    margin: 0;
                                  "
                                >
                                  <span style="font-size: 12px;"
                                    >© Copyright {{.Year}} {{.Company}}</span
                                  >
                                </p>
                              </div>
                            </div>
                            <!--[if mso]></td></tr></table><![endif]-->
                            <!--[if (!mso)&(!IE)]><!-->
                          </div>
                          <!--<![endif]-->
                        </div>
                      </div>
                      <!--[if (mso)|(IE)]></td></tr></table><![endif]-->
                      <!--[if (mso)|(IE)]></td></tr></table></td></tr></table><![endif]-->
                    </div>
                  </div>
                </div>
                <!--[if (mso)|(IE)]></td></tr></table><![endif]-->
              </td>
            </tr>
          </tbody>
        </table>
        <!--[if (IE)]></div><![endif]-->
      </body>
    </html>`

type BrandedEmailData struct {
	Year        string
	Address     string
	Company     string
	HtmlContent string
	LogoLink    string
	Domain      string
	HeaderColor string
	Date        string
}

var emailTemplate *template.Template

func init() {
	emptyTemplate := template.New("Email Template")
	filledTemplate, err := emptyTemplate.Parse(rawEmailTemplate)
	if err != nil {
		panic(err)
	}

	emailTemplate = filledTemplate
}

func FormatDigestEmail(recipient string, HTMLContent string, textContent string) *gomail.Message {
	bodyBuffer := new(bytes.Buffer)

	date := time.Now().Format("02.01.2006")
	emailTemplate.Execute(bodyBuffer, BrandedEmailData{
		Company:     "Baida",
		Address:     "",
		LogoLink:    "",
		Domain:      "",
		HeaderColor: "#000",
		Year:        strconv.Itoa(time.Now().Year()),
		Date:        date,
		HtmlContent: HTMLContent,
	})

	m := gomail.NewMessage()
	m.SetHeader("From", "feeds@baida.dev")
	m.SetHeader("To", recipient)
	m.SetHeader("Subject", "The Weekly Digest - "+date)
	m.SetBody("text/html", bodyBuffer.String())

	if textContent != "" {
		m.AddAlternative("text/plain", textContent)
	}

	return m
}
