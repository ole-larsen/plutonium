import { CommonOptions } from "./common-options";
import { Contract } from "./contract";
export declare const defaults: Required<GovernorOptions>;
export declare const votesOptions: readonly ["erc20votes", "erc721votes", "comp"];
export declare type VotesOptions = typeof votesOptions[number];
export declare const timelockOptions: readonly [false, "openzeppelin", "compound"];
export declare type TimelockOptions = typeof timelockOptions[number];
export declare function printGovernor(opts?: GovernorOptions): string;
export interface GovernorOptions extends CommonOptions {
    name: string;
    delay: string;
    period: string;
    blockTime?: number;
    proposalThreshold?: string;
    decimals?: number;
    quorumMode?: 'percent' | 'absolute';
    quorumPercent?: number;
    quorumAbsolute?: string;
    votes?: VotesOptions;
    timelock?: TimelockOptions;
    bravo?: boolean;
    settings?: boolean;
}
export declare function isAccessControlRequired(opts: Partial<GovernorOptions>): boolean;
export declare function buildGovernor(opts: GovernorOptions): Contract;
export declare const numberPattern: RegExp;
//# sourceMappingURL=governor.d.ts.map