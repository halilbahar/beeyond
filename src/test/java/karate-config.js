function fn(){
    var env = karate.env;
    karate.log('karate.env system property was:', env);
    if (!env) {
        env = 'dev';
    }
    var config = {
        baseUrl: 'http://localhost:8080',
        testTeacherCredentials: '{ username: "testteacher", password: "teacher"}',
        testStudentCredentials: '{ username: "teststudent", password: "student"}'
    };
    karate.configure('connectTimeout', 5000);
    karate.configure('readTimeout', 5000);
    return config;
}