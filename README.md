# summon-aws

AWS provider for [Summon](https://conjurinc.github.io/summon).

Provides access to secrets stored in Amazon S3.

## Usage

[Set summon-aws as your Summon provider](https://github.com/conjurinc/summon#flags).

Give summon a path to an object in S3 and it will fetch it for you and
print the value to stdout.

```bash
summon --provider summon-aws \
--yaml 'MONGOPASS: !var myorg-creds/aws/dev/mongodb-password' \
printenv MONGOPASS
8h9psadf89sdahfp98
```

`myorg-credentials` is the bucket name.

## Configuration

summon-aws uses the [official AWS Go SDK](https://github.com/aws/aws-sdk-go). It will use
the credentials file or environment variables [as they explain](https://github.com/aws/aws-sdk-go#configuring-credentials).