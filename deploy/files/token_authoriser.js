const AWS = require('aws-sdk');

AWS.config.region = process.env.AWS_DEFAULT_REGION || 'eu-west-2';

function getSecret(secretKey) {
    return new Promise((resolve, reject) => {
        const ssm = new AWS.SSM();
        const params = {
            Name: secretKey,
            WithDecryption: true
        };
        console.log("calling sdk ssm");
        ssm.getParameter(params, function (err, data) {
            console.log("ssm - request made");
            if (err) {
                console.error(err);
                return reject(err);
            }
            if (!data.Parameter) {
                console.error("secret not found");
                return reject(new Error('Token not found'));
            }
            console.log("Found secret");
            return resolve(data.Parameter.Value);
        });
    });
}

exports.handler = async (event) => {
    console.log(event)
    let response = {
        "isAuthorized": false,
        "context": {
            "stringKey": "value",
            "numberKey": 1,
            "booleanKey": true,
            "arrayKey": ["value1", "value2"],
            "mapKey": {"value1": "value2"}
        }
    };
    let authHeader = event.headers['authorization']
    console.log("Auth Header", authHeader)
    await getSecret(process.env.SECRET_KEY)
        .then(secret => {
            if (!authHeader === secret) {
                response = {
                    "isAuthorized": false,
                    "context": {
                        "booleanKey": true,
                    }
                }
            } else {
                response = {
                    "isAuthorized": true,
                    "context": {
                        "booleanKey": true,
                    }
                }
            }
        })
        .catch(err => {
            // trap any errors and respond with '500 Server Error'
            console.error(err);
            response = {
                "isAuthorized": false
            }
        });
    //
    // if (event.headers.authorization === "secretToken") {
    //     response = {
    //         "isAuthorized": true,
    //         "context": {
    //             "stringKey": "value",
    //             "numberKey": 1,
    //             "booleanKey": true,
    //             "arrayKey": ["value1", "value2"],
    //             "mapKey": {"value1": "value2"}
    //         }
    //     };
    // }

    console.log("Response", response)
    return response;

};
