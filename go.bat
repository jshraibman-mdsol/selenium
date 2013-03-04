@echo off

java -Xmx2048m -XX:MaxPermSize=1024m -XX:ReservedCodeCacheSize=256m -client -jar third_party\jruby\jruby-complete.jar -r third_party\jruby\gems\albacore.jar -r third_party\jruby\gems\rake.jar -r third_party\jruby\gems\childprocess.jar -X-C -S rake %*
