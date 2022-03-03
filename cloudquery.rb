require 'ffi'
require 'json'

# Module that represents shared lib
module Cloudquery
  extend FFI::Library

  ffi_lib File.dirname(__FILE__) + '/cloudquery.so'

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
                  [String.by_value],
                  :int

  def aws(aws_json)
    json_string = JSON.generate(aws_json)
    c_string = Cloudquery::String.new(json_string)
    Cloudquery.QueryAWS(c_string)
  end

  module_function :aws

end

aws_credentials = {
  aws_access_key_id: "AWS_ACCESS_KEY_ID",
  aws_secret_access_key: "AWS_SECRET_ACCESS_KEY",
  aws_session_token: "AWS_SESSION_TOKEN",
  region: "AWS_REGION",
}

Cloudquery.aws(aws_credentials)
