import { useErrorsStore } from "@/stores/errors"

export const setLoginError = (msg: string) => {
    useErrorsStore().setLoginError(msg)
}

export const setRegisterErrorUsername = (msg: string) => {
    useErrorsStore().setRegisterErrorUsername(msg)
}

export const setRegisterErrorEmail = (msg: string) => {
    useErrorsStore().setRegisterErrorEmail(msg)
}

export const setGroupError = (msg: string) => {
    useErrorsStore().setGroupError(msg)
}
