"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.numberPattern = exports.buildGovernor = exports.isAccessControlRequired = exports.printGovernor = exports.timelockOptions = exports.votesOptions = exports.defaults = void 0;
const common_functions_1 = require("./common-functions");
const common_options_1 = require("./common-options");
const contract_1 = require("./contract");
const error_1 = require("./error");
const set_access_control_1 = require("./set-access-control");
const print_1 = require("./print");
const set_info_1 = require("./set-info");
const set_upgradeable_1 = require("./set-upgradeable");
const define_functions_1 = require("./utils/define-functions");
const duration_1 = require("./utils/duration");
exports.defaults = {
    name: 'MyGovernor',
    delay: '1 block',
    period: '1 week',
    votes: 'erc20votes',
    timelock: 'openzeppelin',
    blockTime: 13.2,
    decimals: 18,
    proposalThreshold: '0',
    quorumMode: 'percent',
    quorumPercent: 4,
    quorumAbsolute: '',
    bravo: false,
    settings: true,
    access: common_options_1.defaults.access,
    upgradeable: common_options_1.defaults.upgradeable,
    info: common_options_1.defaults.info
};
exports.votesOptions = ['erc20votes', 'erc721votes', 'comp'];
exports.timelockOptions = [false, 'openzeppelin', 'compound'];
function printGovernor(opts = exports.defaults) {
    return (0, print_1.printContract)(buildGovernor(opts));
}
exports.printGovernor = printGovernor;
function isAccessControlRequired(opts) {
    return opts.upgradeable === 'uups';
}
exports.isAccessControlRequired = isAccessControlRequired;
function withDefaults(opts) {
    var _a, _b, _c, _d, _e, _f, _g, _h;
    return {
        ...opts,
        ...(0, common_options_1.withCommonDefaults)(opts),
        decimals: (_a = opts.decimals) !== null && _a !== void 0 ? _a : exports.defaults.decimals,
        blockTime: opts.blockTime || exports.defaults.blockTime,
        quorumPercent: (_b = opts.quorumPercent) !== null && _b !== void 0 ? _b : exports.defaults.quorumPercent,
        quorumAbsolute: (_c = opts.quorumAbsolute) !== null && _c !== void 0 ? _c : exports.defaults.quorumAbsolute,
        proposalThreshold: opts.proposalThreshold || exports.defaults.proposalThreshold,
        settings: (_d = opts.settings) !== null && _d !== void 0 ? _d : exports.defaults.settings,
        bravo: (_e = opts.bravo) !== null && _e !== void 0 ? _e : exports.defaults.bravo,
        quorumMode: (_f = opts.quorumMode) !== null && _f !== void 0 ? _f : exports.defaults.quorumMode,
        votes: (_g = opts.votes) !== null && _g !== void 0 ? _g : exports.defaults.votes,
        timelock: (_h = opts.timelock) !== null && _h !== void 0 ? _h : exports.defaults.timelock
    };
}
function buildGovernor(opts) {
    const allOpts = withDefaults(opts);
    const c = new contract_1.ContractBuilder(allOpts.name);
    validateDecimals(allOpts.decimals);
    addBase(c, allOpts);
    addSettings(c, allOpts);
    addCounting(c, allOpts);
    addBravo(c, allOpts);
    addVotes(c, allOpts);
    addQuorum(c, allOpts);
    addTimelock(c, allOpts);
    (0, set_access_control_1.setAccessControl)(c, allOpts.access);
    (0, set_upgradeable_1.setUpgradeable)(c, allOpts.upgradeable, allOpts.access);
    (0, set_info_1.setInfo)(c, allOpts.info);
    return c;
}
exports.buildGovernor = buildGovernor;
function addBase(c, { name }) {
    c.addParent({
        name: 'Governor',
        path: '@openzeppelin/contracts/governance/Governor.sol',
    }, [name]);
    c.addOverride('IGovernor', functions.votingDelay);
    c.addOverride('IGovernor', functions.votingPeriod);
    c.addOverride('IGovernor', functions.quorum);
    c.addOverride('Governor', functions.state);
    c.addOverride('Governor', functions.propose);
    c.addOverride('Governor', functions.proposalThreshold);
    c.addOverride('Governor', functions._execute);
    c.addOverride('Governor', functions._cancel);
    c.addOverride('Governor', functions._executor);
    c.addOverride('Governor', common_functions_1.supportsInterface);
}
function addSettings(c, allOpts) {
    if (allOpts.settings) {
        c.addParent({
            name: 'GovernorSettings',
            path: '@openzeppelin/contracts/governance/extensions/GovernorSettings.sol',
        }, [
            { value: getVotingDelay(allOpts), note: allOpts.delay },
            { value: getVotingPeriod(allOpts), note: allOpts.period },
            { lit: getProposalThreshold(allOpts) },
        ]);
        c.addOverride('GovernorSettings', functions.votingDelay, 'view');
        c.addOverride('GovernorSettings', functions.votingPeriod, 'view');
        c.addOverride('GovernorSettings', functions.proposalThreshold, 'view');
    }
    else {
        setVotingParameters(c, allOpts);
        setProposalThreshold(c, allOpts);
    }
}
function getVotingDelay(opts) {
    try {
        return (0, duration_1.durationToBlocks)(opts.delay, opts.blockTime);
    }
    catch (e) {
        if (e instanceof Error) {
            throw new error_1.OptionsError({
                delay: e.message,
            });
        }
        else {
            throw e;
        }
    }
}
function getVotingPeriod(opts) {
    try {
        return (0, duration_1.durationToBlocks)(opts.period, opts.blockTime);
    }
    catch (e) {
        if (e instanceof Error) {
            throw new error_1.OptionsError({
                period: e.message,
            });
        }
        else {
            throw e;
        }
    }
}
function validateDecimals(decimals) {
    if (!/^\d+$/.test(decimals.toString())) {
        throw new error_1.OptionsError({
            decimals: 'Not a valid number',
        });
    }
}
function getProposalThreshold({ proposalThreshold, decimals, votes }) {
    if (!/^\d+$/.test(proposalThreshold)) {
        throw new error_1.OptionsError({
            proposalThreshold: 'Not a valid number',
        });
    }
    if (/^0+$/.test(proposalThreshold) || decimals === 0 || votes === 'erc721votes') {
        return proposalThreshold;
    }
    else {
        return `${proposalThreshold}e${decimals}`;
    }
}
function setVotingParameters(c, opts) {
    const delayBlocks = getVotingDelay(opts);
    c.setFunctionBody([`return ${delayBlocks}; // ${opts.delay}`], functions.votingDelay);
    const periodBlocks = getVotingPeriod(opts);
    c.setFunctionBody([`return ${periodBlocks}; // ${opts.period}`], functions.votingPeriod);
}
function setProposalThreshold(c, opts) {
    const threshold = getProposalThreshold(opts);
    const nonZeroThreshold = parseInt(threshold) !== 0;
    if (nonZeroThreshold) {
        c.setFunctionBody([`return ${threshold};`], functions.proposalThreshold);
    }
}
function addCounting(c, { bravo }) {
    if (!bravo) {
        c.addParent({
            name: 'GovernorCountingSimple',
            path: '@openzeppelin/contracts/governance/extensions/GovernorCountingSimple.sol',
        });
    }
}
const votesModules = {
    erc20votes: {
        tokenType: 'IVotes',
        parentName: 'GovernorVotes',
    },
    erc721votes: {
        tokenType: 'IVotes',
        parentName: 'GovernorVotes',
    },
    comp: {
        tokenType: 'ERC20VotesComp',
        parentName: 'GovernorVotesComp',
    },
};
function addVotes(c, { votes }) {
    const tokenArg = '_token';
    const { tokenType, parentName } = votesModules[votes];
    c.addConstructorArgument({
        type: tokenType,
        name: tokenArg,
    });
    c.addParent({
        name: parentName,
        path: `@openzeppelin/contracts/governance/extensions/${parentName}.sol`,
    }, [{ lit: tokenArg }]);
}
exports.numberPattern = /^(?!$)(\d*)(?:\.(\d+))?(?:e(\d+))?$/;
function addQuorum(c, opts) {
    if (opts.quorumMode === 'percent') {
        if (opts.votes !== 'erc20votes' && opts.votes !== 'erc721votes') {
            throw new error_1.OptionsError({
                quorumPercent: 'Percent-based quorum is only available for ERC20Votes or ERC721Votes',
            });
        }
        if (opts.quorumPercent > 100) {
            throw new error_1.OptionsError({
                quorumPercent: 'Invalid percentage',
            });
        }
        let { quorumFractionNumerator, quorumFractionDenominator } = getQuorumFractionComponents(opts.quorumPercent);
        if (quorumFractionDenominator !== undefined) {
            c.addOverride('GovernorVotesQuorumFraction', functions.quorumDenominator);
            c.setFunctionBody([
                `return ${quorumFractionDenominator};`
            ], functions.quorumDenominator, 'pure');
        }
        c.addParent({
            name: 'GovernorVotesQuorumFraction',
            path: '@openzeppelin/contracts/governance/extensions/GovernorVotesQuorumFraction.sol',
        }, [quorumFractionNumerator]);
        c.addOverride('GovernorVotesQuorumFraction', functions.quorum);
    }
    else if (opts.quorumMode === 'absolute') {
        if (!exports.numberPattern.test(opts.quorumAbsolute)) {
            throw new error_1.OptionsError({
                quorumAbsolute: 'Not a valid number',
            });
        }
        let returnStatement = (opts.decimals === 0 || opts.votes === 'erc721votes') ?
            `return ${opts.quorumAbsolute};` :
            `return ${opts.quorumAbsolute}e${opts.decimals};`;
        c.setFunctionBody([
            returnStatement,
        ], functions.quorum, 'pure');
    }
}
const timelockModules = {
    openzeppelin: {
        timelockType: 'TimelockController',
        parentName: 'GovernorTimelockControl',
    },
    compound: {
        timelockType: 'ICompoundTimelock',
        parentName: 'GovernorTimelockCompound',
    },
};
function getQuorumFractionComponents(quorumPercent) {
    let quorumFractionNumerator = quorumPercent;
    let quorumFractionDenominator = undefined;
    const quorumPercentSegments = quorumPercent.toString().split(".");
    if (quorumPercentSegments.length > 2) {
        throw new error_1.OptionsError({
            quorumPercent: 'Invalid percentage',
        });
    }
    else if (quorumPercentSegments.length == 2 && quorumPercentSegments[0] !== undefined && quorumPercentSegments[1] !== undefined) {
        quorumFractionNumerator = parseInt(quorumPercentSegments[0].concat(quorumPercentSegments[1]));
        const decimals = quorumPercentSegments[1].length;
        quorumFractionDenominator = '100';
        while (quorumFractionDenominator.length < decimals + 3) {
            quorumFractionDenominator += '0';
        }
    }
    return { quorumFractionNumerator, quorumFractionDenominator };
}
function addTimelock(c, { timelock }) {
    if (timelock === false) {
        return;
    }
    const timelockArg = '_timelock';
    const { timelockType, parentName } = timelockModules[timelock];
    c.addConstructorArgument({
        type: timelockType,
        name: timelockArg,
    });
    c.addParent({
        name: parentName,
        path: `@openzeppelin/contracts/governance/extensions/${parentName}.sol`,
    }, [{ lit: timelockArg }]);
    c.addOverride('IGovernor', functions.propose);
    c.addOverride(parentName, functions._execute);
    c.addOverride(parentName, functions._cancel);
    c.addOverride(parentName, functions._executor);
    c.addOverride(parentName, functions.state);
    c.addOverride(parentName, common_functions_1.supportsInterface);
}
function addBravo(c, { bravo, timelock }) {
    if (bravo) {
        if (timelock === false) {
            throw new error_1.OptionsError({
                timelock: 'GovernorBravo compatibility requires a timelock',
            });
        }
        c.addParent({
            name: 'GovernorCompatibilityBravo',
            path: '@openzeppelin/contracts/governance/compatibility/GovernorCompatibilityBravo.sol',
        });
        c.addOverride('IGovernor', functions.state);
        c.addOverride('GovernorCompatibilityBravo', functions.propose);
        c.addOverride('IERC165', common_functions_1.supportsInterface);
    }
}
const functions = (0, define_functions_1.defineFunctions)({
    votingDelay: {
        args: [],
        returns: ['uint256'],
        kind: 'public',
        mutability: 'pure',
    },
    votingPeriod: {
        args: [],
        returns: ['uint256'],
        kind: 'public',
        mutability: 'pure',
    },
    proposalThreshold: {
        args: [],
        returns: ['uint256'],
        kind: 'public',
        mutability: 'pure',
    },
    quorum: {
        args: [
            { name: 'blockNumber', type: 'uint256' },
        ],
        returns: ['uint256'],
        kind: 'public',
        mutability: 'view',
    },
    quorumDenominator: {
        args: [],
        returns: ['uint256'],
        kind: 'public',
        mutability: 'view',
    },
    propose: {
        args: [
            { name: 'targets', type: 'address[] memory' },
            { name: 'values', type: 'uint256[] memory' },
            { name: 'calldatas', type: 'bytes[] memory' },
            { name: 'description', type: 'string memory' },
        ],
        returns: ['uint256'],
        kind: 'public',
    },
    _execute: {
        args: [
            { name: 'proposalId', type: 'uint256' },
            { name: 'targets', type: 'address[] memory' },
            { name: 'values', type: 'uint256[] memory' },
            { name: 'calldatas', type: 'bytes[] memory' },
            { name: 'descriptionHash', type: 'bytes32' },
        ],
        kind: 'internal',
    },
    _cancel: {
        args: [
            { name: 'targets', type: 'address[] memory' },
            { name: 'values', type: 'uint256[] memory' },
            { name: 'calldatas', type: 'bytes[] memory' },
            { name: 'descriptionHash', type: 'bytes32' },
        ],
        returns: ['uint256'],
        kind: 'internal',
    },
    state: {
        args: [
            { name: 'proposalId', type: 'uint256' },
        ],
        returns: ['ProposalState'],
        kind: 'public',
        mutability: 'view',
    },
    _executor: {
        args: [],
        returns: ['address'],
        kind: 'internal',
        mutability: 'view',
    },
});
//# sourceMappingURL=governor.js.map