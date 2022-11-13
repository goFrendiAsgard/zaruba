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

    getResponseErrorMessage(response, errorMessage='Error') {
        const errorSuffix = response.status ? ` (HTTP Status: ${response.status})` : ' (No HTTP Status)';
        if (response && response.data && typeof(response.data) == 'string') {
            return `${errorMessage} ${errorSuffix}:\n- ${response.data}`;
        }
        if (response && response.data && response.data.detail && typeof(response.data.detail) == 'string') {
            return `${errorMessage} ${errorSuffix}:\n- ${response.data.detail}`;
        }
        if (response && response.data && response.data.detail && Array.isArray(response.data.detail)) {
            const details = response.data.detail;
            const detailErrorMessageList = [];
            for(const detail of details) {
                detailErrorMessageList.push(this._getDetailErrorMessage(detail));
            }
            const detailErrorMessage = detailErrorMessageList.join('\n');
            return `${errorMessage} ${errorSuffix}:\n${detailErrorMessage}`;
        }
        return `${errorMessage} ${errorSuffix}`;
    }

    _getDetailErrorMessage(detail) {
        const msgStr = this._getMsgStr(detail.msg, detail.type);
        const locStr = this._getLocStr(detail.loc);
        if (locStr && locStr != '') {
            return `- ${msgStr}: ${locStr}`;
        }
        return `- ${msgStr}`;
    }

    _getLocStr(loc) {
        if (Array.isArray(loc)) {
            if (loc.length == 0) {
                return '';
            }
            if (loc.length > 1 && loc[0] == 'body') {
                return loc.slice(1).join('.');
            }
            return loc.join('.');
        }
        return loc;
    }

    _getMsgStr(msg, msgType) {
        if (msgType && msgType != '') {
            return `(${msgType}) ${msg}`;
        }
        return msg;
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
            handleModule(type, source, path, options) {
                if (type === '.json') {
                    return JSON.parse(source);
                }
            }
        }
    }

}