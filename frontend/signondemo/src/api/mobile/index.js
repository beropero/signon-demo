import { req } from '@/utils/request'

export function mobileSignin(form){
    return req({
        method: 'post',
        url: '/signon/user/phone/login',
        data: {
            mobile: form.account,
            password: form.pass,
            verification_code: form.code,
        },
    })
}

export function mobileSendCode(form) {
    return req({
        method: 'post',
        url: '/signon/user/phone/sendsms',
        data: {
            mobile: form.account
        }
    })
}

export function mobileSignUp(form) {
    return req({
        method: 'post',
        url: '/signon/user/phone/signup',
        data: {
            mobile: form.account,
            password: form.pass1,
            password2: form.pass2,
            verification_code: form.code,
        }
    })
}