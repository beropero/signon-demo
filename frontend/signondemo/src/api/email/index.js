import { req } from '@/utils/request'

export function emailSignin(form) {
    return req({
        method: 'post',
        url: '/signon/user/email/login',
        data: {
            email: form.account,
            password: form.pass,
            verification_code: form.code,
        }
    })
}

export function emailSendCode(form) {
    return req({
        method: 'post',
        url: '/signon/user/email/sendcode',
        data: {
            email: form.account
        }
    })
}

export function emailSignUp(form) {
    return req({
        method: 'post',
        url: '/signon/user/email/signup',
        data: {
            email: form.account,
            password: form.pass1,
            password2: form.pass2,
            verification_code: form.code,
        }
    })
}