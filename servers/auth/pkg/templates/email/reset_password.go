package templates

import (
	"bytes"
	"text/template"
)

var ResetPasswordTemplate string

type ResetPasswordTemplateData struct {
	Username   string
	ResetLink  string
	ExpireTime int
}

func GenerateResetPasswordTemplate(data ResetPasswordTemplateData) string {
	ResetPasswordTemplate = `<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd"><html xmlns="http://www.w3.org/1999/xhtml"><head>
  <meta name="viewport" content="width=device-width, initial-scale=1.0"/>
  <meta http-equiv="Content-Type" content="text/html; charset=UTF-8"/>
  
<style type="text/css">*:not(br):not(tr):not(html) {
font-family: Arial, 'Helvetica Neue', Helvetica, sans-serif !important;
-webkit-box-sizing: border-box !important;
box-sizing: border-box !important
}cite:before {
content: "\2014 \0020" !important
}@media only screen and (max-width: 600px){
.email-body_inner,
      .email-footer {
width: 100% !important
}
}
@media only screen and (max-width: 500px){
.button {
width: 100% !important
}
}
</style></head>
<body dir="ltr" style="height:100%;margin:0;line-height:1.4;background-color:#F2F4F6;color:#74787E;-webkit-text-size-adjust:none;width:100%">
  <table class="email-wrapper" width="100%" cellpadding="0" cellspacing="0" style="width:100%;margin:0;padding:0;background-color:#F2F4F6">
    <tbody><tr>
      <td class="content" style="color:#74787E;font-size:15px;line-height:18px;align:center;padding:0">
        <table class="email-content" width="100%" cellpadding="0" cellspacing="0" style="width:100%;margin:0;padding:0">
          
          <tbody><tr>
            <td class="email-masthead" style="color:#74787E;font-size:15px;line-height:18px;padding:25px 0;text-align:center">
              <a class="email-masthead_name" href="https://beta.truongng.me" target="_blank" style="font-size:16px;font-weight:bold;color:#2F3133;text-decoration:none;text-shadow:0 1px 0 white">
                
                  <img src="https://i.imgur.com/g0jbCNf.png" class="email-logo" style="max-height:50px"/>
                
                </a>
            </td>
          </tr>

          
          <tr>
            <td class="email-body" width="100%" style="color:#74787E;font-size:15px;line-height:18px;width:100%;margin:0;padding:0;border-top:1px solid #EDEFF2;border-bottom:1px solid #EDEFF2;background-color:#FFF">
              <table class="email-body_inner" align="center" width="570" cellpadding="0" cellspacing="0" style="width:570px;margin:0 auto;padding:0">
                
                <tbody><tr>
                  <td class="content-cell" style="color:#74787E;font-size:15px;line-height:18px;padding:35px">
                    <h1 style="margin-top:0;color:#2F3133;font-size:19px;font-weight:bold">Dear {{.Username}},</h1>
                    
                        
                          
                            <p style="margin-top:0;color:#74787E;font-size:16px;line-height:1.5em">You have received this email because a password reset request for your account was received.</p>
                          
                        
                    
                    

                      

                      
                      
                        
                        
                        
                      

                      
                      
                        
                          
                            <p style="margin-top:0;color:#74787E;font-size:16px;line-height:1.5em">To reset your password, please click here. This link will expire in {{.ExpireTime}} minutes.</p>
                            
                            
                            
                              <!--[if mso]>
                              
                                <div style="margin: 30px auto;v-text-anchor:middle;text-align:center">
                                  <v:roundrect xmlns:v="urn:schemas-microsoft-com:vml" 
                                    xmlns:w="urn:schemas-microsoft-com:office:word" 
                                    href="{{.ResetLink}}" 
                                    style="height:45px;v-text-anchor:middle;width:200px;background-color:#ce5734;"
                                    arcsize="10%" 
                                    strokecolor="#ce5734" fillcolor="#ce5734"
                                    >
                                    <w:anchorlock/>
                                    <center style="color: #FFFFFF;font-size: 15px;text-align: center;font-family:sans-serif;font-weight:bold;">
                                      Reset your password
                                    </center>
                                  </v:roundrect>
                                </div>
                              
                                 
                              <![endif]-->
                              <!--[if !mso]><!-- -->
                              <table class="body-action" align="center" width="100%" cellpadding="0" cellspacing="0" style="width:100%;margin:30px auto;padding:0;text-align:center">
                                <tbody><tr>
                                  <td align="center" style="padding:10px 5px;color:#74787E;font-size:15px;line-height:18px">
                                    <div>
                                      
                                        <a href="{{.ResetLink}}" class="button" style="display:inline-block;border-radius:3px;font-size:15px;line-height:45px;text-align:center;text-decoration:none;-webkit-text-size-adjust:none;mso-hide:all;color:#ffffff;background-color:#ce5734;width:200px" target="_blank" width="200">
                                          Reset your password
                                        </a>
                                      
                                      
                                    </div>
                                  </td>
                                </tr>
                              </tbody></table>
                              <!--[endif]---->
                          
                            <p style="margin-top:0;color:#74787E;font-size:16px;line-height:1.5em">If you did not request a password reset, no further action is required on your part.</p>
                            
                            
                            
                              <!--[if mso]>
                              
                                 
                              <![endif]-->
                              <!--[if !mso]><!-- -->
                              <table class="body-action" align="center" width="100%" cellpadding="0" cellspacing="0" style="width:100%;margin:30px auto;padding:0;text-align:center">
                                <tbody><tr>
                                  <td align="center" style="padding:10px 5px;color:#74787E;font-size:15px;line-height:18px">
                                    <div>
                                      
                                      
                                    </div>
                                  </td>
                                </tr>
                              </tbody></table>
                              <!--[endif]---->
                          
                        
                      

                    
                     
                        
                          
                            <p style="margin-top:0;color:#74787E;font-size:16px;line-height:1.5em">Need help, or have questions? Just reply to this email, we&#39;d love to help.</p>
                          
                        
                      

                    <p style="margin-top:0;color:#74787E;font-size:16px;line-height:1.5em">
                      Sincerely,
                      <br/>
                      The poker
                    </p>

                    
                       
                        <table class="body-sub" style="width:100%;margin-top:25px;padding-top:25px;border-top:1px solid #EDEFF2;table-layout:fixed">
                          <tbody>
                              
                                
                                <tr>
                                  <td style="padding:10px 5px;color:#74787E;font-size:15px;line-height:18px">
                                    <p class="sub" style="margin-top:0;color:#74787E;line-height:1.5em;font-size:12px">If the verify button is not working for you, just copy and paste the URL below into your web browser.</p>
                                    <p class="sub" style="margin-top:0;color:#74787E;line-height:1.5em;font-size:12px"><a href="{{.ResetLink}}" style="color:#3869D4;word-break:break-all">{{.ResetLink}}</a></p>
                                  </td>
                                </tr>
                                
                              
                                
                              
                          </tbody>
                        </table>
                      
                    
                  </td>
                </tr>
              </tbody></table>
            </td>
          </tr>
          <tr>
            <td style="padding:10px 5px;color:#74787E;font-size:15px;line-height:18px">
              <table class="email-footer" align="center" width="570" cellpadding="0" cellspacing="0" style="width:570px;margin:0 auto;padding:0;text-align:center">
                <tbody><tr>
                  <td class="content-cell" style="color:#74787E;font-size:15px;line-height:18px;padding:35px">
                    <p class="sub center" style="margin-top:0;line-height:1.5em;color:#AEAEAE;font-size:12px;text-align:center">
                      Copyright © 2023 The Poker. All rights reserved.
                    </p>
                  </td>
                </tr>
              </tbody></table>
            </td>
          </tr>
        </tbody></table>
      </td>
    </tr>
  </tbody></table>


</body></html>`

	tmpl, err := template.New("ResetPasswordTemplate").Parse(ResetPasswordTemplate)
	if err != nil {
		panic(err)
	}
	buf := &bytes.Buffer{}
	err = tmpl.Execute(buf, data)
	if err != nil {
		panic(err)
	}

	return buf.String()
}
