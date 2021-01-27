function fn(){
    var env = karate.env;
    karate.log('karate.env system property was:', env);
    if (!env) {
        env = 'dev';
    }

    var basicAuth = function(isTeacher) {
        var temp;
        if (isTeacher){
            temp = 'stuetz:password';
        } else{
            temp = 'moritz:password';
        }
        var Base64 = Java.type('java.util.Base64');
        var encoded = Base64.getEncoder().encodeToString(temp.bytes);
        return 'Basic ' + encoded;
    };
    var config = {
        baseUrl: 'http://localhost:8081',
        auth: basicAuth(),
        config: karate.tags
    };
    karate.configure('connectTimeout', 5000);
    karate.configure('readTimeout', 5000);
    return config;
}
