require 'ffi'

# Module that represents shared lib
module Cloudquery
  extend FFI::Library

  ffi_lib './cloudquery.so'

  # define class String to map:
  # C type struct { const char *p; GoInt n; }
  # There's some bug exchanging string values between rails and go
  class String < FFI::Struct
    layout :p,     :pointer,
           :len,   :long_long
    def initialize(str)
      self[:p] = FFI::MemoryPointer.from_string(str)
      self[:len] = str.bytesize
      self
    end
  end

  # foreign function definitions
  attach_function :QueryAWS,
                  [String.by_value, String.by_value],
                  :int
end

# Call QueryAWS (credentials, config)
puts "Launching Go:",  Cloudquery.QueryAWS(
  Cloudquery::String.new("[default]\naws_access_key_id = RubyCredentials\naws_secret_access_key = RubyCredentials\n"),
  Cloudquery::String.new("[default]\nregion = us-west-2\noutput=json\n")
)
