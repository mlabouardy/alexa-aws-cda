# AWS Developer Quiz Skill

[![Become AWS Certified Developer](https://img.youtube.com/vi/9Uty9vVEtOU/0.jpg)](https://www.youtube.com/watch?v=9Uty9vVEtOU)


# How it works

<p>
  <img src="quiz.png" />
</p>

# IAM Policy

```
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Sid": "1",
            "Effect": "Allow",
            "Action": [
                "logs:CreateLogStream",
                "logs:CreateLogGroup",
                "logs:PutLogEvents"
            ],
            "Resource": "*"
        },
        {
            "Sid": "2",
            "Effect": "Allow",
            "Action": "dynamodb:Scan",
            "Resource": [
                "arn:aws:dynamodb:us-east-1:ACCOUNT_ID:table/Questions/index/ID",
                "arn:aws:dynamodb:us-east-1:ACCOUNT_ID:table/Questions"
            ]
        }
    ]
}
```

# Maintainers

Mohamed Labouardy

# License

MIT
