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
  # Returns a json serializable string
  # Input String: {"success": true, "message": "success", "logs": "Logs here"}
  # Output Object: "{\"success\":true,\"message\":\"success\",\"logs\":\"Logs here\"}"
  
  # ----------------- AWS -----------------
  attach_function :QueryAWS,
                  [String.by_value, String.by_value],
                  :strptr

  # Returns [Boolean, String, String]
  def aws(aws_json, database)
    json_string = JSON.generate(aws_json)
    c_string = Cloudquery::String.new(json_string)
    db_c_string = Cloudquery::String.new(database)
    output, pointer = Cloudquery.QueryAWS(c_string, db_c_string)
    output
  end

  # ----------------- GCP -----------------
  attach_function :QueryGCP,
                  [String.by_value, String.by_value],
                  :strptr

  # Returns [Boolean, String, String]
  def gcp(json_string, database)
    c_string = Cloudquery::String.new(json_string)
    db_c_string = Cloudquery::String.new(database)
    output, pointer = Cloudquery.QueryGCP(c_string, db_c_string)
    output
  end

  # ----------------- Azure -----------------
  attach_function :QueryAzure,
                  [String.by_value, String.by_value],
                  :strptr

  # Returns [Boolean, String, String]
  def azure(azure_json, database)
    json_string = JSON.generate(azure_json)
    c_string = Cloudquery::String.new(json_string)
    db_c_string = Cloudquery::String.new(database)
    output, pointer = Cloudquery.QueryAzure(c_string, db_c_string)
    output
  end

  # ----------------- Digitalocean -----------------
  attach_function :QueryDigitalocean,
                  [String.by_value, String.by_value],
                  :strptr

  # Returns [Boolean, String, String]
  def digitalocean(digitalocean_json, database)
    json_string = JSON.generate(digitalocean_json)
    c_string = Cloudquery::String.new(json_string)
    db_c_string = Cloudquery::String.new(database)
    output, pointer = Cloudquery.QueryDigitalocean(c_string, db_c_string)
    output
  end

  # ----------------- Kubernetes -----------------
  attach_function :QueryKubernetes,
                  [String.by_value, String.by_value],
                  :strptr

  # Returns [Boolean, String, String]
  def kubernetes(json_string, database)
    c_string = Cloudquery::String.new(json_string)
    db_c_string = Cloudquery::String.new(database)
    output, pointer = Cloudquery.QueryKubernetes(c_string, db_c_string)
    output
  end

  # ----------------- Okta -----------------
  attach_function :QueryOkta,
                  [String.by_value, String.by_value],
                  :strptr

  # Returns [Boolean, String, String]
  def okta(okta_json, database)
    json_string = JSON.generate(okta_json)
    c_string = Cloudquery::String.new(json_string)
    db_c_string = Cloudquery::String.new(database)
    output, pointer = Cloudquery.QueryOkta(c_string, db_c_string)
    output
  end

  # ---------------------------------------------------
  
  # module functions
  # Each function returns an array: [successBoolean, messageString]
  # success: true for success, false for fail
  # message: "Description for success or fail"
  module_function :aws, :gcp, :azure, :digitalocean, :kubernetes, :okta

end

