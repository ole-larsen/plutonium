import 'array.prototype.flatmap/auto';
export declare type Lines = string | typeof whitespace | Lines[];
declare const whitespace: unique symbol;
export declare function formatLines(...lines: Lines[]): string;
export declare function spaceBetween(...lines: Lines[][]): Lines[];
export {};
//# sourceMappingURL=format-lines.d.ts.map