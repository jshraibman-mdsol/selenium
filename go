#!/bin/bash

# we want jruby-complete to take care of all things ruby
unset GEM_HOME
unset GEM_PATH

# we don't trust `-S rake`, which will search PATH, so pass -e to JRuby
java -client \
     -Xmx2048m \
     -XX:MaxPermSize=1024m \
     -XX:ReservedCodeCacheSize=256m \
     -jar third_party/jruby/jruby-complete.jar \
     $(for jar in ./third_party/jruby/gems/*.jar; do echo -r $jar; done) \
     -X-C \
     -rubygems \
     -e 'require "rake"; Rake.application.run' \
     -- $*

