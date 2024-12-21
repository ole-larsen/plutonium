export declare type OptionsErrorMessages = {
    [prop in string]?: string;
};
export declare class OptionsError extends Error {
    readonly messages: OptionsErrorMessages;
    constructor(messages: OptionsErrorMessages);
}
//# sourceMappingURL=error.d.ts.map