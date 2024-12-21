import type { BaseFunction } from '../contract';
declare type ImplicitNameFunction = Omit<BaseFunction, 'name'>;
export declare function defineFunctions<F extends string>(fns: Record<F, ImplicitNameFunction>): Record<F, BaseFunction>;
export {};
//# sourceMappingURL=define-functions.d.ts.map