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
  # Returns a json serialisable string
  # Input String: "{\"success\":true,\"message\":\"success\"}"
  # Output Object: {"success": true, "message": "success"}
  
  # ----------------- AWS -----------------
  attach_function :QueryAWS,
                  [String.by_value, String.by_value],
                  :strptr

  # Returns [Boolean, String]
  def aws(aws_json, database)
    json_string = JSON.generate(aws_json)
    c_string = Cloudquery::String.new(json_string)
    db_c_string = Cloudquery::String.new(database)
    output, pointer = Cloudquery.QueryAWS(c_string, db_c_string)

    data = JSON.parse(output)
    [data["success"], data["message"]]
  end

  # ----------------- GCP -----------------
  attach_function :QueryGCP,
                  [String.by_value, String.by_value],
                  :strptr

  # Returns [Boolean, String]
  def gcp(json_string, database)
    c_string = Cloudquery::String.new(json_string)
    db_c_string = Cloudquery::String.new(database)
    output, pointer = Cloudquery.QueryGCP(c_string, db_c_string)
    
    data = JSON.parse(output)
    [data["success"], data["message"]]
  end

  # ----------------- Azure -----------------
  attach_function :QueryAzure,
                  [String.by_value, String.by_value],
                  :strptr

  # Returns [Boolean, String]
  def azure(azure_json, database)
    json_string = JSON.generate(azure_json)
    c_string = Cloudquery::String.new(json_string)
    db_c_string = Cloudquery::String.new(database)
    output, pointer = Cloudquery.QueryAzure(c_string, db_c_string)

    data = JSON.parse(output)
    [data["success"], data["message"]]
  end

  # ----------------- Digitalocean -----------------
  attach_function :QueryDigitalocean,
                  [String.by_value, String.by_value],
                  :strptr

  # Returns [Boolean, String]
  def digitalocean(digitalocean_json, database)
    json_string = JSON.generate(digitalocean_json)
    c_string = Cloudquery::String.new(json_string)
    db_c_string = Cloudquery::String.new(database)
    output, pointer = Cloudquery.QueryDigitalocean(c_string, db_c_string)

    data = JSON.parse(output)
    [data["success"], data["message"]]
  end

  # ----------------- Kubernetes -----------------
  attach_function :QueryKubernetes,
                  [String.by_value, String.by_value],
                  :strptr

  # Returns [Boolean, String]
  def kubernetes(json_string, database)
    c_string = Cloudquery::String.new(json_string)
    db_c_string = Cloudquery::String.new(database)
    output, pointer = Cloudquery.QueryKubernetes(c_string, db_c_string)

    data = JSON.parse(output)
    [data["success"], data["message"]]
  end

  # ----------------- Okta -----------------
  attach_function :QueryOkta,
                  [String.by_value, String.by_value],
                  :strptr

  # Returns [Boolean, String]
  def okta(okta_json, database)
    json_string = JSON.generate(okta_json)
    c_string = Cloudquery::String.new(json_string)
    db_c_string = Cloudquery::String.new(database)
    output, pointer = Cloudquery.QueryOkta(c_string, db_c_string)

    data = JSON.parse(output)
    [data["success"], data["message"]]
  end

  # ---------------------------------------------------
  
  # module functions
  # Each function returns an array: [successBoolean, messageString]
  # success: true for success, false for fail
  # message: "Description for success or fail"
  module_function :aws, :gcp, :azure, :digitalocean, :kubernetes, :okta

  database = "tsdb://postgres:pass@localhost:5432/cloudtry?sslmode=disable"
  credentials = {
      aws_access_key_id: "AKIATE67RE6LE34ODKPN",
      aws_secret_access_key: "lVUV9tdOO++aIX9IBDNXlAQfVDs5jGh2QNUlOo0",
      region: "us-west-2",
  }

  success, pointer = aws(credentials, database)

  puts success
  puts pointer

end

