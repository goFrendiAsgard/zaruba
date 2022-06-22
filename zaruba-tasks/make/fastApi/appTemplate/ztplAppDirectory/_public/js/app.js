class AppHelper {

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

    async login(createAccessTokenUrl, username, password) {
        const response = await axios.post(createAccessTokenUrl, {username, password});
        if (response && response.status == 200 && response.data && response.data.access_token) {
            this.setCookie('app_access_token', response.data.access_token);
            return true
        }
        if (response && response.data && response.data.detail && response.data.detail.msg) {
            throw new Error(response.data.detail.msg);
        }
        throw new Error('Login failed')
    }

    logout() {
        this.unsetCookie('app_access_token')
    }

    async renewToken(renewAccessTokenUrl, interval) {
        const access_token = this.getCookie('app_access_token')
        if (access_token != '') {
            const response = await axios.post(renewAccessTokenUrl, {access_token});
            if (response && response.status == 200 && response.data && response.data.access_token) {
                this.setCookie('app_access_token', response.data.access_token);
            }
        }
        setTimeout(() => this.renewToken(renewAccessTokenUrl, interval), interval * 1000)
    }

}