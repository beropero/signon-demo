import { req } from "@/utils/request";

export function getQrcode(){
    return req({
        method: 'get',
        url: '/signon/user/wechat/getqrcode',
    })
}

export function checklogin(ticket){
    return req({
        method: 'post',
        url: '/signon/user/wechat/checklogin',
        data: {
            ticket: ticket,
        }
    })
}