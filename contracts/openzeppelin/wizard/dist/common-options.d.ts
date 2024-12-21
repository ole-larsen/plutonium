import type { Access } from "./set-access-control";
import type { Info } from "./set-info";
import type { Upgradeable } from "./set-upgradeable";
export declare const defaults: Required<CommonOptions>;
export interface CommonOptions {
    access?: Access;
    upgradeable?: Upgradeable;
    info?: Info;
}
export declare function withCommonDefaults(opts: CommonOptions): Required<CommonOptions>;
//# sourceMappingURL=common-options.d.ts.map