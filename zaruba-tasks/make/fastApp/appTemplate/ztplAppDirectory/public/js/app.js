class AppHelper {

    constructor(configs) {
        this.axios = configs.axios;
        this.createAccessTokenUrl = configs.createAccessTokenUrl;
        this.renewAccessTokenUrl = configs.renewAccessTokenUrl;
        this.renewAccessTokenInterval = configs.renewAccessTokenInterval;
        this.accessTokenCookieKey = configs.accessTokenCookieKey;
        this.appMode = configs.appMode;
    }

    setCookie(cookieName, cookieValue) {
        document.cookie = cookieName + '=' + cookieValue + ';path=/';
    }

    unsetCookie(cookieName) {
        document.cookie = cookieName + '=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/'
    }

    getCookie(cookieName) {
        let valuePrefix = cookieName + '=';
        let cookieParts = document.cookie.split(';');
        for(let partIndex = 0; partIndex < cookieParts.length; partIndex++) {
            let cookiePart = cookieParts[partIndex].trim();
            if (cookiePart.indexOf(valuePrefix) == 0) {
                return cookiePart.substring(valuePrefix.length, cookiePart.length);
            }
        }
        return '';
    }

    getResponseErrorMessage(response, defaultErrorMessage='error') {
        const errorSuffix = response.status ? ` (HTTP Status: ${response.status})` : ' (No HTTP Status)';
        if (response && response.data && typeof response.data == 'string') {
            return response.data + errorSuffix;
        }
        if (response && response.data && response.data.detail && typeof response.data.detail == 'string') {
            return response.data.detail + errorSuffix;
        }
        if (response && response.data && response.data.detail && response.data.detail.msg && typeof response.data.detail.msg == 'string') {
            return response.data.detail.msg + errorSuffix;
        }
        return defaultErrorMessage + errorSuffix;
    }

    async login(username, password) {
        const response = await this.axios.post(this.createAccessTokenUrl, {username, password});
        if (response && response.status == 200 && response.data && response.data.access_token) {
            this.setCookie(this.accessTokenCookieKey, response.data.access_token);
            return true
        }
        throw new Error(this.getResponseErrorMessage(response, 'Login failed'));
    }

    logout() {
        this.unsetCookie(this.accessTokenCookieKey)
    }

    getAuthBearer() {
        const access_token = this.getCookie(this.accessTokenCookieKey)
        if (access_token != '') {
            return 'Bearer ' + access_token;
        }
        return '';
    }

    getConfigAuthHeader() {
        const authBearer = this.getAuthBearer();
        if (authBearer != '') {
            return {headers: {'Authorization': authBearer}};
        }
        return {};
    }

    async renewAccessToken() {
        const access_token = this.getCookie(this.accessTokenCookieKey)
        if (access_token != '') {
            try {
                const response = await this.axios.post(this.renewAccessTokenUrl, {access_token}, this.getConfigAuthHeader());
                if (response && response.status == 200 && response.data && response.data.access_token) {
                    this.setCookie(this.accessTokenCookieKey, response.data.access_token);
                } else {
                    this.unsetCookie(this.accessTokenCookieKey);
                }
            } catch(error) {
                console.error(error);
            }
        }
        setTimeout(() => this.renewAccessToken(this.renewAccessTokenUrl, this.renewAccessTokenInterval), this.renewAccessTokenInterval * 1000)
    }

    alert(message) {
        return window.alert(message);
    }

    confirm(message) {
        return window.confirm(message);
    }

    alertError(error) {
        console.error(error);
        this.alert(error.message);
    }

    getVueSfcLoaderOptions() {
        return {
            moduleCache: {
                vue: Vue,
            },
            async getFile(url) {
                const res = await fetch(url);
                if ( !res.ok ) {
                    throw Object.assign(new Error(res.statusText + ' ' + url), { res });
                }
                return res.text();
            },
            addStyle(textContent) {
                const style = Object.assign(document.createElement('style'), { textContent });
                const ref = document.head.getElementsByTagName('style')[0] || null;
                document.head.insertBefore(style, ref);
            },
            log(type, ...args) {
                console[type](...args);
            },
            compiledCache: {
                set(key, str) {
                    // naive storage space management
                    for (; ;) {
                        try {
                            // doc: https://developer.mozilla.org/en-US/docs/Web/API/Storage
                            window.localStorage.setItem(key, str);
                            break;
                        } catch (ex) {
                            // handle: Uncaught DOMException: Failed to execute 'setItem' on 'Storage': Setting the value of 'XXX' exceeded the quota
                            window.localStorage.removeItem(window.localStorage.key(0));
                        }
                    }
                },
                get(key) {
                    return window.localStorage.getItem(key);
                },
            },
            handleModule(type, source, path, options) {
                if (type === '.json') {
                    return JSON.parse(source);
                }
            }
        }
    }

}