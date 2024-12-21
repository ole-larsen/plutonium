import { GenericOptions } from '../build-generic';
import type { Contract } from '../contract';
declare type Subset = 'all' | 'minimal-cover';
export declare function generateOptions(): Generator<GenericOptions>;
interface GeneratedContract {
    id: string;
    options: GenericOptions;
    contract: Contract;
}
interface GeneratedSource extends GeneratedContract {
    source: string;
}
export declare function generateSources(subset: Subset): Generator<GeneratedSource>;
export declare function writeGeneratedSources(dir: string, subset: Subset): Promise<void>;
export {};
//# sourceMappingURL=sources.d.ts.map