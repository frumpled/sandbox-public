
# Docs / References
- [Blog Post](https://andrewgriffithsonline.com/blog/180412-migrate-go-api-to-serverless-under-10-mins/#convert-application-code-serverless) that gives real example and calls out libs to simplify the effort.
- [AWS Docs for acceptable Lambda Handler signatures](https://docs.aws.amazon.com/lambda/latest/dg/golang-handler.html)

# Notes

This was a real bitch to get working -\_-
- Had to use special libs to wrap the handler given by the Twirp generated code to inject into a lambda handler
	- Can simplify this if AWS Lambda ever let's you stick a plain old http handler into their lambdas
- Had to use API Gateway to make this a public endpoint that can be hit
- Had to rewrite the path to exclude the API Gateway "stage name"

# Wins
- DID get this working
- Easy to swap out the generated Twirp handlers into more lambdas
- Twirp client works w/o issues and w/o snowflake configs

