import type { ContractBuilder } from './contract';
import { Access } from './set-access-control';
export declare const upgradeableOptions: readonly [false, "transparent", "uups"];
export declare type Upgradeable = typeof upgradeableOptions[number];
export declare function setUpgradeable(c: ContractBuilder, upgradeable: Upgradeable, access: Access): void;
//# sourceMappingURL=set-upgradeable.d.ts.map