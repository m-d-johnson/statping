// DO NOT EDIT ** This file was generated with go generate on 2020-08-06 04:13:25.078006 +0000 UTC ** DO NOT EDIT //
package notifiers

const emailSuccess = `<!doctype html><html xmlns=http://www.w3.org/1999/xhtml xmlns:v=urn:schemas-microsoft-com:vml xmlns:o=urn:schemas-microsoft-com:office:office><title>Statping Service Notification</title><meta http-equiv=x-ua-compatible content="IE=edge"><meta charset=utf-8><meta name=viewport content="width=device-width,initial-scale=1"><style type=text/css>
          #outlook a { padding:0; }
          body { margin:0;padding:0;-webkit-text-size-adjust:100%;-ms-text-size-adjust:100%; }
          table, td { border-collapse:collapse;mso-table-lspace:0pt;mso-table-rspace:0pt; }
          img { border:0;height:auto;line-height:100%; outline:none;text-decoration:none;-ms-interpolation-mode:bicubic; }
          p { display:block;margin:13px 0; }
        </style><link href="https://fonts.googleapis.com/css?family=Ubuntu:300,400,500,700" rel=stylesheet type=text/css><style type=text/css>
          @import url(https://fonts.googleapis.com/css?family=Ubuntu:300,400,500,700);
        </style><style type=text/css>
      @media only screen and (min-width:480px) {
        .mj-column-per-100 { width:100% !important; max-width: 100%; }
      }
    </style><style type=text/css>
        
        

    @media only screen and (max-width:480px) {
      table.mj-full-width-mobile { width: 100% !important; }
      td.mj-full-width-mobile { width: auto !important; }
    }
  
        </style><body style=background-color:#E7E7E7;><div style=background-color:#E7E7E7;><div style="background:#12ab0c url(https://assets.statping.com/greenbackground.png) top center / auto repeat;margin:0px auto;max-width:600px;"><div style=line-height:0;font-size:0;><table align=center background=https://assets.statping.com/greenbackground.png border=0 cellpadding=0 cellspacing=0 role=presentation style="background:#12ab0c url(https://assets.statping.com/greenbackground.png) top center / auto repeat;width:100%;"><tbody><tr><td style=direction:ltr;font-size:0px;padding:0px;text-align:center;><div class="mj-column-per-100 mj-outlook-group-fix" style=font-size:0px;text-align:left;direction:ltr;display:inline-block;vertical-align:top;width:100%;><table border=0 cellpadding=0 cellspacing=0 role=presentation style=vertical-align:top; width=100%><tr><td align=center style="font-size:0px;padding:10px 25px;word-break:break-word;"><table border=0 cellpadding=0 cellspacing=0 role=presentation style=border-collapse:collapse;border-spacing:0px;><tbody><tr><td style=width:45px;><a href=https://statping.com target=_blank><img alt=Sphero height=auto src=https://assets.statping.com/iconlight.png style=border:0;display:block;outline:none;text-decoration:none;height:auto;width:100%;font-size:13px; width=45></a></table></table></div></table></div></div><div style="background:#ffffff;background-color:#ffffff;margin:0px auto;max-width:600px;"><table align=center border=0 cellpadding=0 cellspacing=0 role=presentation style=background:#ffffff;background-color:#ffffff;width:100%;><tbody><tr><td style="direction:ltr;font-size:0px;padding:20px 0;text-align:center;"><div class="mj-column-per-100 mj-outlook-group-fix" style=font-size:0px;text-align:left;direction:ltr;display:inline-block;vertical-align:top;width:100%;><table border=0 cellpadding=0 cellspacing=0 role=presentation style=vertical-align:top; width=100%><tr><td align=left style=font-size:0px;padding:15px;word-break:break-word;><div style="font-family:Ubuntu, Helvetica, Arial, sans-serif;font-size:22px;line-height:30px;text-align:left;color:#000000;">{{.Service.Name}} is currently offline, you might want to check it.</div><tr><td style="font-size:0px;padding:20px 0;padding-top:10px;padding-right:0px;padding-bottom:10px;padding-left:0px;word-break:break-word;"><div style="margin:0px auto;max-width:600px;"><table align=center border=0 cellpadding=0 cellspacing=0 role=presentation style=width:100%;><tbody><tr><td style="direction:ltr;font-size:0px;padding:20px 0;padding-bottom:10px;padding-left:0px;padding-right:0px;padding-top:10px;text-align:center;"><div class="mj-column-per-100 mj-outlook-group-fix" style=font-size:0px;text-align:left;direction:ltr;display:inline-block;vertical-align:top;width:100%;><table border=0 cellpadding=0 cellspacing=0 role=presentation style=vertical-align:top; width=100%><tr><td align=center style="font-size:0px;padding:10px 25px;word-break:break-word;"><div style="font-family:Ubuntu, Helvetica, Arial, sans-serif;font-size:20px;line-height:1;text-align:center;color:#626262;">Offline for {{.Service.Downtime.Human}}</div><tr><td align=center vertical-align=middle style="font-size:0px;padding:10px 25px;word-break:break-word;"><table border=0 cellpadding=0 cellspacing=0 role=presentation style=border-collapse:separate;line-height:100%;><tr><td align=center bgcolor=#4caf50 role=presentation style="border:none;border-radius:4px;cursor:auto;mso-padding-alt:10px 25px;background:#4caf50;" valign=middle><a href={{.Core.Domain}}/service/{{.Service.Id}} style="display:inline-block;background:#4caf50;color:#ffffff;font-family:Ubuntu, Helvetica, Arial, sans-serif;font-size:13px;font-weight:normal;line-height:120%;margin:0;text-decoration:none;text-transform:none;padding:10px 25px;mso-padding-alt:0px;border-radius:4px;" target=_blank>View Dashboard</a></table></table></div></table></div><tr><td style=font-size:0px;padding:0px;word-break:break-word;><div style="background:#fafafa;background-color:#fafafa;margin:0px auto;max-width:600px;"><table align=center border=0 cellpadding=0 cellspacing=0 role=presentation style=background:#fafafa;background-color:#fafafa;width:100%;><tbody><tr><td style=direction:ltr;font-size:0px;padding:0px;text-align:center;><div class="mj-column-per-100 mj-outlook-group-fix" style=font-size:0px;text-align:left;direction:ltr;display:inline-block;vertical-align:top;width:100%;><table border=0 cellpadding=0 cellspacing=0 role=presentation style=vertical-align:top; width=100%><tr><td align=left style="font-size:0px;padding:10px 25px;word-break:break-word;"><div style="font-family:Ubuntu, Helvetica, Arial, sans-serif;font-size:20px;line-height:1;text-align:left;color:#626262;">Service Domain</div><tr><td align=left style="font-size:0px;padding:10px 25px;padding-top:0px;word-break:break-word;"><div style="font-family:Ubuntu, Helvetica, Arial, sans-serif;font-size:14px;line-height:1;text-align:left;color:#626262;">{{.Service.Domain}}</div></table></div></table></div><tr><td style=font-size:0px;padding:0px;word-break:break-word;><div style="background:#ffffff;background-color:#ffffff;margin:0px auto;max-width:600px;"><table align=center border=0 cellpadding=0 cellspacing=0 role=presentation style=background:#ffffff;background-color:#ffffff;width:100%;><tbody><tr><td style=direction:ltr;font-size:0px;padding:0px;text-align:center;><div class="mj-column-per-100 mj-outlook-group-fix" style=font-size:0px;text-align:left;direction:ltr;display:inline-block;vertical-align:top;width:100%;><table border=0 cellpadding=0 cellspacing=0 role=presentation style=vertical-align:top; width=100%><tr><td align=left style="font-size:0px;padding:10px 25px;word-break:break-word;"><div style="font-family:Ubuntu, Helvetica, Arial, sans-serif;font-size:20px;line-height:1;text-align:left;color:#626262;">Current Issue</div><tr><td align=left style="font-size:0px;padding:10px 25px;padding-top:0px;word-break:break-word;"><div style="font-family:Ubuntu, Helvetica, Arial, sans-serif;font-size:14px;line-height:1;text-align:left;color:#626262;">{{.Failure.Issue}}</div></table></div></table></div><tr><td style=font-size:0px;word-break:break-word;><div style=height:30px;>&nbsp;</div><tr><td style=font-size:0px;padding:0;word-break:break-word;><div style="background:#12ab0c url(https://assets.statping.com/greenbackground.png) top center / auto repeat;margin:0px auto;max-width:600px;"><div style=line-height:0;font-size:0;><table align=center background=https://assets.statping.com/greenbackground.png border=0 cellpadding=0 cellspacing=0 role=presentation style="background:#12ab0c url(https://assets.statping.com/greenbackground.png) top center / auto repeat;width:100%;"><tbody><tr><td style=direction:ltr;font-size:0px;padding:0;text-align:center;><div class="mj-column-per-100 mj-outlook-group-fix" style=font-size:0px;text-align:left;direction:ltr;display:inline-block;vertical-align:top;width:100%;><table border=0 cellpadding=0 cellspacing=0 role=presentation style=vertical-align:top; width=100%><tr><td align=center style="font-size:0px;padding:10px 25px;word-break:break-word;"><table border=0 cellpadding=0 cellspacing=0 role=presentation style=border-collapse:collapse;border-spacing:0px;><tbody><tr><td style=width:250px;><a href=https://www.sphero.com target=_blank><img height=auto src=https://assets.statping.com/statpingcom.png style=border:0;display:block;outline:none;text-decoration:none;height:auto;width:100%;font-size:13px; width=250></a></table></table></div></table></div></div><tr><td style="font-size:0px;padding:20px 0;padding-top:10px;padding-bottom:0;word-break:break-word;"><div style="margin:0px auto;max-width:600px;"><table align=center border=0 cellpadding=0 cellspacing=0 role=presentation style=width:100%;><tbody><tr><td style="direction:ltr;font-size:0px;padding:20px 0;padding-bottom:0;padding-top:10px;text-align:center;"><div class="mj-column-per-100 mj-outlook-group-fix" style=font-size:0px;text-align:left;direction:ltr;display:inline-block;vertical-align:top;width:100%;><table border=0 cellpadding=0 cellspacing=0 role=presentation style=vertical-align:top; width=100%><tr><td align=center style="font-size:0px;padding:10px 25px;word-break:break-word;"><div style="font-family:Ubuntu, Helvetica, Arial, sans-serif;font-size:11px;line-height:16px;text-align:center;color:#445566;">You are receiving this email because one of your services has changed on your Statping instance. You can modify this email on the Email Notifier page in Settings.</div><tr><td align=center style="font-size:0px;padding:10px 25px;word-break:break-word;"><div style="font-family:Ubuntu, Helvetica, Arial, sans-serif;font-size:11px;line-height:16px;text-align:center;color:#445566;">&copy; Statping</div></table></div></table></div><tr><td style="font-size:0px;padding:20px 0;padding-top:0;padding-bottom:0;word-break:break-word;"><div style="margin:0px auto;max-width:600px;"><table align=center border=0 cellpadding=0 cellspacing=0 role=presentation style=width:100%;><tbody><tr><td style="direction:ltr;font-size:0px;padding:20px 0;padding-bottom:0;padding-top:0;text-align:center;"><div class="mj-column-per-100 mj-outlook-group-fix" style=font-size:0;line-height:0;text-align:left;display:inline-block;width:100%;direction:ltr;><div class="mj-column-per-100 mj-outlook-group-fix" style=font-size:0px;text-align:left;direction:ltr;display:inline-block;vertical-align:top;width:100%;><table border=0 cellpadding=0 cellspacing=0 role=presentation width=100%><tbody><tr><td style=vertical-align:top;padding-right:0;><table border=0 cellpadding=0 cellspacing=0 role=presentation width=100%><tr><td align=center style="font-size:0px;padding:10px 25px;word-break:break-word;"><div style="font-family:Ubuntu, Helvetica, Arial, sans-serif;font-size:11px;font-weight:bold;line-height:16px;text-align:center;color:#445566;"><a class=footer-link href=https://statping.com>Statping.com</a>&#160;&#160;&#160;&#160;&#160;&#160;&#160;&#160;
<a class=footer-link href=https://github.com/statping/statping>Github</a>&#160;&#160;&#160;&#160;&#160;&#160;&#160;&#160;
<a class=footer-link href=https://statping.com/privacy>Privacy</a>&#160;&#160;&#160;&#160;&#160;&#160;&#160;&#160;
<a class=footer-link href="https://news.statping.com/delete?email={{.Email}}">Unsubscribe</a></div></table></table></div></div></table></div></table></div></table></div></div>`

const emailFailure = `<!doctype html><html xmlns=http://www.w3.org/1999/xhtml xmlns:v=urn:schemas-microsoft-com:vml xmlns:o=urn:schemas-microsoft-com:office:office><title>Statping Service Notification</title><meta http-equiv=x-ua-compatible content="IE=edge"><meta charset=utf-8><meta name=viewport content="width=device-width,initial-scale=1"><style type=text/css>
          #outlook a { padding:0; }
          body { margin:0;padding:0;-webkit-text-size-adjust:100%;-ms-text-size-adjust:100%; }
          table, td { border-collapse:collapse;mso-table-lspace:0pt;mso-table-rspace:0pt; }
          img { border:0;height:auto;line-height:100%; outline:none;text-decoration:none;-ms-interpolation-mode:bicubic; }
          p { display:block;margin:13px 0; }
        </style><link href="https://fonts.googleapis.com/css?family=Ubuntu:300,400,500,700" rel=stylesheet type=text/css><style type=text/css>
          @import url(https://fonts.googleapis.com/css?family=Ubuntu:300,400,500,700);
        </style><style type=text/css>
      @media only screen and (min-width:480px) {
        .mj-column-per-100 { width:100% !important; max-width: 100%; }
      }
    </style><style type=text/css>
        
        

    @media only screen and (max-width:480px) {
      table.mj-full-width-mobile { width: 100% !important; }
      td.mj-full-width-mobile { width: auto !important; }
    }
  
        </style><body style=background-color:#E7E7E7;><div style=background-color:#E7E7E7;><div style="background:#a30911 url(https://assets.statping.com/offlinebanner.png) top center / auto repeat;margin:0px auto;max-width:600px;"><div style=line-height:0;font-size:0;><table align=center background=https://assets.statping.com/offlinebanner.png border=0 cellpadding=0 cellspacing=0 role=presentation style="background:#a30911 url(https://assets.statping.com/offlinebanner.png) top center / auto repeat;width:100%;"><tbody><tr><td style=direction:ltr;font-size:0px;padding:0px;text-align:center;><div class="mj-column-per-100 mj-outlook-group-fix" style=font-size:0px;text-align:left;direction:ltr;display:inline-block;vertical-align:top;width:100%;><table border=0 cellpadding=0 cellspacing=0 role=presentation style=vertical-align:top; width=100%><tr><td align=center style="font-size:0px;padding:10px 25px;word-break:break-word;"><table border=0 cellpadding=0 cellspacing=0 role=presentation style=border-collapse:collapse;border-spacing:0px;><tbody><tr><td style=width:45px;><a href=https://statping.com target=_blank><img alt=Sphero height=auto src=https://assets.statping.com/iconlight.png style=border:0;display:block;outline:none;text-decoration:none;height:auto;width:100%;font-size:13px; width=45></a></table></table></div></table></div></div><div style="background:#ffffff;background-color:#ffffff;margin:0px auto;max-width:600px;"><table align=center border=0 cellpadding=0 cellspacing=0 role=presentation style=background:#ffffff;background-color:#ffffff;width:100%;><tbody><tr><td style="direction:ltr;font-size:0px;padding:20px 0;text-align:center;"><div class="mj-column-per-100 mj-outlook-group-fix" style=font-size:0px;text-align:left;direction:ltr;display:inline-block;vertical-align:top;width:100%;><table border=0 cellpadding=0 cellspacing=0 role=presentation style=vertical-align:top; width=100%><tr><td align=left style=font-size:0px;padding:15px;word-break:break-word;><div style="font-family:Ubuntu, Helvetica, Arial, sans-serif;font-size:22px;line-height:30px;text-align:left;color:#000000;">{{.Service.Name}} is currently offline, you might want to check it.</div><tr><td style="font-size:0px;padding:20px 0;padding-top:10px;padding-right:0px;padding-bottom:10px;padding-left:0px;word-break:break-word;"><div style="margin:0px auto;max-width:600px;"><table align=center border=0 cellpadding=0 cellspacing=0 role=presentation style=width:100%;><tbody><tr><td style="direction:ltr;font-size:0px;padding:20px 0;padding-bottom:10px;padding-left:0px;padding-right:0px;padding-top:10px;text-align:center;"><div class="mj-column-per-100 mj-outlook-group-fix" style=font-size:0px;text-align:left;direction:ltr;display:inline-block;vertical-align:top;width:100%;><table border=0 cellpadding=0 cellspacing=0 role=presentation style=vertical-align:top; width=100%><tr><td align=center style="font-size:0px;padding:10px 25px;word-break:break-word;"><div style="font-family:Ubuntu, Helvetica, Arial, sans-serif;font-size:20px;line-height:1;text-align:center;color:#626262;">Offline for {{.Service.Downtime.Human}}</div><tr><td align=center vertical-align=middle style="font-size:0px;padding:10px 25px;word-break:break-word;"><table border=0 cellpadding=0 cellspacing=0 role=presentation style=border-collapse:separate;line-height:100%;><tr><td align=center bgcolor=#cb121c role=presentation style="border:none;border-radius:4px;cursor:auto;mso-padding-alt:10px 25px;background:#cb121c;" valign=middle><a href={{.Core.Domain}}/service/{{.Service.Id}} style="display:inline-block;background:#cb121c;color:#ffffff;font-family:Ubuntu, Helvetica, Arial, sans-serif;font-size:13px;font-weight:normal;line-height:120%;margin:0;text-decoration:none;text-transform:none;padding:10px 25px;mso-padding-alt:0px;border-radius:4px;" target=_blank>View Dashboard</a></table></table></div></table></div><tr><td style=font-size:0px;padding:0px;word-break:break-word;><div style="background:#fafafa;background-color:#fafafa;margin:0px auto;max-width:600px;"><table align=center border=0 cellpadding=0 cellspacing=0 role=presentation style=background:#fafafa;background-color:#fafafa;width:100%;><tbody><tr><td style=direction:ltr;font-size:0px;padding:0px;text-align:center;><div class="mj-column-per-100 mj-outlook-group-fix" style=font-size:0px;text-align:left;direction:ltr;display:inline-block;vertical-align:top;width:100%;><table border=0 cellpadding=0 cellspacing=0 role=presentation style=vertical-align:top; width=100%><tr><td align=left style="font-size:0px;padding:10px 25px;word-break:break-word;"><div style="font-family:Ubuntu, Helvetica, Arial, sans-serif;font-size:20px;line-height:1;text-align:left;color:#626262;">Service Domain</div><tr><td align=left style="font-size:0px;padding:10px 25px;padding-top:0px;word-break:break-word;"><div style="font-family:Ubuntu, Helvetica, Arial, sans-serif;font-size:14px;line-height:1;text-align:left;color:#626262;">{{.Service.Domain}}</div></table></div></table></div><tr><td style=font-size:0px;padding:0px;word-break:break-word;><div style="background:#ffffff;background-color:#ffffff;margin:0px auto;max-width:600px;"><table align=center border=0 cellpadding=0 cellspacing=0 role=presentation style=background:#ffffff;background-color:#ffffff;width:100%;><tbody><tr><td style=direction:ltr;font-size:0px;padding:0px;text-align:center;><div class="mj-column-per-100 mj-outlook-group-fix" style=font-size:0px;text-align:left;direction:ltr;display:inline-block;vertical-align:top;width:100%;><table border=0 cellpadding=0 cellspacing=0 role=presentation style=vertical-align:top; width=100%><tr><td align=left style="font-size:0px;padding:10px 25px;word-break:break-word;"><div style="font-family:Ubuntu, Helvetica, Arial, sans-serif;font-size:20px;line-height:1;text-align:left;color:#626262;">Current Issue</div><tr><td align=left style="font-size:0px;padding:10px 25px;padding-top:0px;word-break:break-word;"><div style="font-family:Ubuntu, Helvetica, Arial, sans-serif;font-size:14px;line-height:1;text-align:left;color:#626262;">{{.Failure.Issue}}</div></table></div></table></div><tr><td style=font-size:0px;word-break:break-word;><div style=height:30px;>&nbsp;</div><tr><td style=font-size:0px;padding:0;word-break:break-word;><div style="background:#a30911 url(https://assets.statping.com/offlinebanner.png) top center / auto repeat;margin:0px auto;max-width:600px;"><div style=line-height:0;font-size:0;><table align=center background=https://assets.statping.com/offlinebanner.png border=0 cellpadding=0 cellspacing=0 role=presentation style="background:#a30911 url(https://assets.statping.com/offlinebanner.png) top center / auto repeat;width:100%;"><tbody><tr><td style=direction:ltr;font-size:0px;padding:0;text-align:center;><div class="mj-column-per-100 mj-outlook-group-fix" style=font-size:0px;text-align:left;direction:ltr;display:inline-block;vertical-align:top;width:100%;><table border=0 cellpadding=0 cellspacing=0 role=presentation style=vertical-align:top; width=100%><tr><td align=center style="font-size:0px;padding:10px 25px;word-break:break-word;"><table border=0 cellpadding=0 cellspacing=0 role=presentation style=border-collapse:collapse;border-spacing:0px;><tbody><tr><td style=width:250px;><a href=https://www.sphero.com target=_blank><img height=auto src=https://assets.statping.com/statpingcom.png style=border:0;display:block;outline:none;text-decoration:none;height:auto;width:100%;font-size:13px; width=250></a></table></table></div></table></div></div><tr><td style="font-size:0px;padding:20px 0;padding-top:10px;padding-bottom:0;word-break:break-word;"><div style="margin:0px auto;max-width:600px;"><table align=center border=0 cellpadding=0 cellspacing=0 role=presentation style=width:100%;><tbody><tr><td style="direction:ltr;font-size:0px;padding:20px 0;padding-bottom:0;padding-top:10px;text-align:center;"><div class="mj-column-per-100 mj-outlook-group-fix" style=font-size:0px;text-align:left;direction:ltr;display:inline-block;vertical-align:top;width:100%;><table border=0 cellpadding=0 cellspacing=0 role=presentation style=vertical-align:top; width=100%><tr><td align=center style="font-size:0px;padding:10px 25px;word-break:break-word;"><div style="font-family:Ubuntu, Helvetica, Arial, sans-serif;font-size:11px;line-height:16px;text-align:center;color:#445566;">You are receiving this email because one of your services has changed on your Statping instance. You can modify this email on the Email Notifier page in Settings.</div><tr><td align=center style="font-size:0px;padding:10px 25px;word-break:break-word;"><div style="font-family:Ubuntu, Helvetica, Arial, sans-serif;font-size:11px;line-height:16px;text-align:center;color:#445566;">&copy; Statping</div></table></div></table></div><tr><td style="font-size:0px;padding:20px 0;padding-top:0;padding-bottom:0;word-break:break-word;"><div style="margin:0px auto;max-width:600px;"><table align=center border=0 cellpadding=0 cellspacing=0 role=presentation style=width:100%;><tbody><tr><td style="direction:ltr;font-size:0px;padding:20px 0;padding-bottom:0;padding-top:0;text-align:center;"><div class="mj-column-per-100 mj-outlook-group-fix" style=font-size:0;line-height:0;text-align:left;display:inline-block;width:100%;direction:ltr;><div class="mj-column-per-100 mj-outlook-group-fix" style=font-size:0px;text-align:left;direction:ltr;display:inline-block;vertical-align:top;width:100%;><table border=0 cellpadding=0 cellspacing=0 role=presentation width=100%><tbody><tr><td style=vertical-align:top;padding-right:0;><table border=0 cellpadding=0 cellspacing=0 role=presentation width=100%><tr><td align=center style="font-size:0px;padding:10px 25px;word-break:break-word;"><div style="font-family:Ubuntu, Helvetica, Arial, sans-serif;font-size:11px;font-weight:bold;line-height:16px;text-align:center;color:#445566;"><a class=footer-link href=https://statping.com>Statping.com</a>&#160;&#160;&#160;&#160;&#160;&#160;&#160;&#160;
<a class=footer-link href=https://github.com/statping/statping>Github</a>&#160;&#160;&#160;&#160;&#160;&#160;&#160;&#160;
<a class=footer-link href=https://statping.com/privacy>Privacy</a>&#160;&#160;&#160;&#160;&#160;&#160;&#160;&#160;
<a class=footer-link href="https://news.statping.com/delete?email={{.Email}}">Unsubscribe</a></div></table></table></div></div></table></div></table></div></table></div></div>`
