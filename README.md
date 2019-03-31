# smart-ads-gcp
Backend for a Ads system in Go with Tools in GCP, including GCS, BigTable, Docker, GKE and etc.

To setup gcloud ssh with emacs, please add the following snippet into init.el (for emacs) or into dotspacemacs/user-init for spacemacs.
```lisp
(require 'tramp)
(add-to-list 'tramp-methods
             '("gcssh"
               (tramp-login-program        "gcloud compute ssh")
               (tramp-login-args           (("%h")))
               (tramp-async-args           (("-q")))
               (tramp-remote-shell         "/bin/sh")
               (tramp-remote-shell-args    ("-c"))
               (tramp-gw-args              (("-o" "GlobalKnownHostsFile=/dev/null")
                                            ("-o" "UserKnownHostsFile=/dev/null")
                                            ("-o" "StrictHostKeyChecking=no")))
               (tramp-default-port         22)))
```
With this, gcould compute VM can be accessed with the following command:
~SPC f f~
~/gcssh:<username>@<instance-name>:<filename>~ (~filename~ is optional).
