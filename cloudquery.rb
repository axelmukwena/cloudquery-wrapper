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
  
  # ----------------- AWS -----------------
  attach_function :QueryAWS,
                  [String.by_value],
                  :int

  def aws(aws_json)
    json_string = JSON.generate(aws_json)
    c_string = Cloudquery::String.new(json_string)
    Cloudquery.QueryAWS(c_string)
  end

  # ----------------- GCP -----------------
  attach_function :QueryGCP,
                  [String.by_value],
                  :int

  def gcp(json_string)
    c_string = Cloudquery::String.new(json_string)
    Cloudquery.QueryGCP(c_string)
  end

  # ----------------- Azure -----------------
  attach_function :QueryAzure,
                  [String.by_value],
                  :int

  def azure(azure_json)
    json_string = JSON.generate(azure_json)
    c_string = Cloudquery::String.new(json_string)
    Cloudquery.QueryAzure(c_string)
  end

  # ----------------- Digitalocean -----------------
  attach_function :QueryDigitalocean,
                  [String.by_value],
                  :int

  def digitalocean(digitalocean_json)
    json_string = JSON.generate(digitalocean_json)
    c_string = Cloudquery::String.new(json_string)
    Cloudquery.QueryDigitalocean(c_string)
  end

  # ----------------- Kubernetes -----------------
  attach_function :QueryKubernetes,
                  [String.by_value],
                  :int

  def kubernetes(json_string)
    c_string = Cloudquery::String.new(json_string)
    Cloudquery.QueryKubernetes(c_string)
  end

  # ----------------- Okta -----------------
  attach_function :QueryOkta,
                  [String.by_value],
                  :int

  def okta(okta_json)
    json_string = JSON.generate(okta_json)
    c_string = Cloudquery::String.new(json_string)
    Cloudquery.QueryOkta(c_string)
  end

  # ---------------------------------------------------
  
  # module functions
  module_function :aws, :gcp, :azure, :digitalocean, :kubernetes, :okta

end