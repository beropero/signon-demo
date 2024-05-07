import { req } from "@/utils/request";
import { useTokenStore } from "@/stores/token";

export function getuserinfo(){
    return req({
        method: 'get',
        url: '/user/getuserinfo',
        headers: {
            'Authorization': useTokenStore().token,
        }
    })
}