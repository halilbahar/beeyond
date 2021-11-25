function fn() {
    let env = karate.env;
    if (!env) {
        env = 'dev';
    }

    let generateString = function () {
        let result = '';
        let characters = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789';
        let charactersLength = characters.length;
        for (let i = 0; i < 300; i++) {
            result += characters.charAt(Math.floor(Math.random() * charactersLength));
        }
        return result;
    };

    let auth = function (tags) {
        let isTeacher = false;
        let isStudent = false;
        for (let i = 0; i < tags.length; i++) {
            let tag = tags[i];
            if (tag == 'teacher') {
                isTeacher = true;
            } else if (tag == 'student') {
                isStudent = true;
            }
        }

        if (isTeacher && isStudent) {
            return 'Please use either @teacher or @student';
        }

        let temp;
        if (isTeacher) {
            temp = 'stuetz:password';
        } else if (isStudent) {
            temp = 'moritz:password';
        }

        let Base64 = Java.type('java.util.Base64');
        let encoded = Base64.getEncoder().encodeToString(Array.from(temp, (x) => x.charCodeAt(0)));
        return 'Basic ' + encoded;
    };

    let config = {
        baseUrl: 'http://localhost:8081',
        auth: auth,
        config: karate.tags,
        generateString: generateString
    };

    let DatabaseCleanup = Java.type('at.htl.beeyond.integration.util.DatabaseCleanup');
    DatabaseCleanup.cleanUp();

    karate.configure('connectTimeout', 5000);
    karate.configure('readTimeout', 5000);
    return config;
}
