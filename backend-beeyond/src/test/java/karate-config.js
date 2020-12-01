function fn(){
    var env = karate.env;
    karate.log('karate.env system property was:', env);
    if (!env) {
        env = 'dev';
    }
    var config = {
        baseUrl: 'http://localhost:8081',
        teacherAuth: function() {
            var temp = 'stuetz:password';
            var Base64 = Java.type('java.util.Base64');
            var encoded = Base64.getEncoder().encodeToString(temp.bytes);
            return 'Basic ' + encoded;
        },
        studentAuth: function() {
            var temp = 'moritz:password';
            var Base64 = Java.type('java.util.Base64');
            var encoded = Base64.getEncoder().encodeToString(temp.bytes);
            return 'Basic ' + encoded;
        }
    };
    karate.configure('connectTimeout', 5000);
    karate.configure('readTimeout', 5000);
    return config;
}
