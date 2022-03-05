package main

import "C"
import (
	"cloudquery/providers"
)

// Main provider functions exported to Ruby

//export QueryAWS
func QueryAWS(awsString string) int {
	ifSuccess := providers.AWS(awsString)
	return ifSuccess
}

//export QueryGCP
func QueryGCP(gcpString string) int {
	ifSuccess := providers.GCP(gcpString)
	return ifSuccess
}

func main() {
	//myString := `{ "type": "service_account", "project_id": "wagestack", "private_key_id": "82f6d09efb316fd6163630b1aa1719e68f44caa2", "private_key": "-----BEGIN PRIVATE KEY-----\nMIIEvAIBADANBgkqhkiG9w0BAQEFAASCBKYwggSiAgEAAoIBAQC6mV+IwUX6GOCY\nUE5YjHT0Jxgqt/oO2gJ6SFbtOioUNwsXxk+BsMAV5j+D0ff57l1I9GvM+o8/jf++\n4YKuTLiu4tzYKMdOh0LGxtQA0Alu+CqcsY1pNGVqNk9f/aaZtgO3WxRkdnQUzLHT\nNB9ofOTjYIIQhoYMqzmMU3geEfpTj2mEV7NevdPG9RzV4Zqi2VIGN71u6yZozNGU\n16Fp6anbQSOrI8lo0aom5rSZLLUD/m+nb7cdMVl6/B1+GB0WO7I96kooduPZLNZm\nJ3RT892O13O2h3jMxXdpgUCVMkpeMJ8YNv0U+rHDgCzivuqjQllQBhg6rw9yrgWG\nagknl4TrAgMBAAECggEABKzXEewVvsBk0Cwi6mEKhRt9pYRahYi8yyeI1gTBDSSb\n6IqVb2BoT6QnLzxRhDeOqsdE9CvK7wK1x6iKx6cwcWFJuzi6VaNZ7vUIVsTfl6Is\nHyTrsBkZ+WFG99a5F9zRlwR3LFUzb1n5kpsZtlp6uZ+vYo6cSTD5DiUBO2+j3/gK\n1iOgaL/oQCsrltTOMaYBzCknwZ1NDkprSSNIU0gCaQcWq5d/X7tzygxdVuGJ2L8p\n6dXJVC1UpiHKcDWNkuWvRHrrO1nztU1CKwUMz3vcrIgKkL3yFels4O+thgWzBUX7\naIbWdIUBw4NJ4VMTV9RE2B4igriD55KvuMXHQRHrUQKBgQD5rQ/VoGjGRruoVESi\ntcXQDoyJQAg4OSghUNh4pRGqtUDOdiuee9GkqQrMLhJIywEmPB/viUjcXpa34M0j\nEhG5TEsRhmekETYeozouOzkNQh/8wSHq5dvk33GATHHECEq8ZETuvXy9z/3aDq9u\nFLSppNIkJbB4B1pP8we6eW2L1wKBgQC/U0+TeQ9csKVhfEAzGaJnv7SSOzG3DvwQ\nD9XFhh3kK7Etz+pMGSYjMSDqGFXdrcMvsdQelG3REqlInZslmaON9zSgHQq38jhI\nFp15xk5HCjBHIPiaFcLIvRzGkTbaD+Un2LOKO4M1qcG3cu/joGMspXAJJut5KNF7\nCfOb4M6NDQKBgHPYYKh2LScSWq/XqaD1RjsrBPoJw8aSfpQ2troDnRbf0pn5KnP2\nb2c/J8tk9Qbhaj8bVpYF1NCq8rOOkp/bGm4ngA05l40Aj2PXyH7665XDQKQ92Ebt\nMAIZysgEsCSM1GBlBbbgJKjNgLNUbQFeihTMbNRoyGBoyPafhM542OMxAoGARaBh\n8z85MfgvF10KWA5aJfuEETttijrvzECXAT0fn6uu3QcvMuZsFJ6KZebZSMU1pSPI\nGCDYHh/2bzC8B2D0PnPaOPKYtfx2MvXX9TsPvZadnyUGk7ybmEYKNNEf7xedw3R/\nUiz6QQs4LjSrzGDP9q12Kj55rywFoAstFmsnf/kCgYADQYdNX5jGlC70NJa0+Tn1\ndT/wghJkBvZczI4/XWG6irrcXAqef+VWkPLm7cAYsiqqRSOe5ObOrsZ9jd2JZO4U\nKgpUZgmmIyRAQvjsAYEtlkinRH2rnfimpyoJ0GPz36QkGMO8Tg0gMFU10P7WLKdw\nye7hanK8EB8sbjmd1cnanw==\n-----END PRIVATE KEY-----\n", "client_email": "wagestack@wagestack.iam.gserviceaccount.com", "client_id": "102364514713233449200", "auth_uri": "https://accounts.google.com/o/oauth2/auth", "token_uri": "https://oauth2.googleapis.com/token", "auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs", "client_x509_cert_url": "https://www.googleapis.com/robot/v1/metadata/x509/wagestack%40wagestack.iam.gserviceaccount.com" }`
	//ifSuccess := providers.GCP(myString)
	//fmt.Printf("Returned: %v", ifSuccess)
}
