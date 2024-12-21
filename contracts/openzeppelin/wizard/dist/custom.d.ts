import { Contract } from './contract';
import { CommonOptions } from './common-options';
export interface CustomOptions extends CommonOptions {
    name: string;
    pausable?: boolean;
}
export declare const defaults: Required<CustomOptions>;
export declare function printCustom(opts?: CustomOptions): string;
export declare function isAccessControlRequired(opts: Partial<CustomOptions>): boolean;
export declare function buildCustom(opts: CustomOptions): Contract;
//# sourceMappingURL=custom.d.ts.map