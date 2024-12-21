import type { ContractBuilder } from "./contract";
export declare const TAG_SECURITY_CONTACT = "@custom:security-contact";
export declare const infoOptions: readonly [{}, {
    readonly securityContact: "security@example.com";
    readonly license: "WTFPL";
}];
export declare const defaults: Info;
export declare type Info = {
    securityContact?: string;
    license?: string;
};
export declare function setInfo(c: ContractBuilder, info: Info): void;
//# sourceMappingURL=set-info.d.ts.map