declare type Blueprint = Record<string, readonly unknown[]>;
declare type Alternatives<B extends Blueprint> = {
    [k in keyof B]: B[k][number];
};
export declare function generateAlternatives<B extends Blueprint>(blueprint: B): Generator<Alternatives<B>>;
export {};
//# sourceMappingURL=alternatives.d.ts.map