function fn() {
    var env = karate.env;
    if (!env) {
        env = 'dev';
    }

    var auth = function (tags) {
        var isTeacher = false;
        var isStudent = false;
        for (var i = 0; i < tags.length; i++) {
            var tag = tags[i];
            if (tag == 'teacher') {
                isTeacher = true;
            } else if (tag == 'student') {
                isStudent = true;
            }
        }

        if (isTeacher && isStudent) {
            return 'Please use either @teacher or @student';
        }

        var temp;
        if (isTeacher) {
            temp = 'stuetz:password';
        } else if (isStudent) {
            temp = 'moritz:password';
        }

        var Base64 = Java.type('java.util.Base64');
        var encoded = Base64.getEncoder().encodeToString(temp.bytes);
        return 'Basic ' + encoded;
    };

    var config = {
        baseUrl: 'http://localhost:8081',
        auth: auth,
        config: karate.tags
    };

    var DatabaseCleanup = Java.type('at.htl.beeyond.integration.util.DatabaseCleanup');
    DatabaseCleanup.cleanUp();

    karate.configure('connectTimeout', 5000);
    karate.configure('readTimeout', 5000);
    return config;
}
