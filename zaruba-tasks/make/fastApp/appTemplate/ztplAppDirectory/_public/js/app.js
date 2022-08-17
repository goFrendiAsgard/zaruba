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

    async login(createAccessTokenUrl, username, password) {
        const response = await axios.post(createAccessTokenUrl, {username, password});
        if (response && response.status == 200 && response.data && response.data.access_token) {
            this.setCookie('app_access_token', response.data.access_token);
            return true
        }
        throw new Error(this.getResponseErrorMessage(response, 'Login failed'));
    }

    logout() {
        this.unsetCookie('app_access_token')
    }

    getAuthBearer() {
        const access_token = this.getCookie('app_access_token')
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

    async renewToken(renewAccessTokenUrl, interval) {
        const access_token = this.getCookie('app_access_token')
        if (access_token != '') {
            try {
                const response = await axios.post(renewAccessTokenUrl, {access_token}, this.getConfigAuthHeader());
                if (response && response.status == 200 && response.data && response.data.access_token) {
                    this.setCookie('app_access_token', response.data.access_token);
                } else {
                    this.unsetCookie('app_access_token');
                }
            } catch(error) {
                console.error(error);
            }
        }
        setTimeout(() => this.renewToken(renewAccessTokenUrl, interval), interval * 1000)
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

}