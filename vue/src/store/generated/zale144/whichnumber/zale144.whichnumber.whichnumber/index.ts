import { txClient, queryClient, MissingWalletError , registry} from './module'

import { EventNewGame } from "./module/types/whichnumber/event"
import { EventRevealTimeout } from "./module/types/whichnumber/event"
import { EventGameEnd } from "./module/types/whichnumber/event"
import { Winner } from "./module/types/whichnumber/event"
import { Params } from "./module/types/whichnumber/params"
import { SystemInfo } from "./module/types/whichnumber/system_info"
import { Game } from "./module/types/whichnumber/types"
import { NumberCommit } from "./module/types/whichnumber/types"
import { NumberReveal } from "./module/types/whichnumber/types"


export { EventNewGame, EventRevealTimeout, EventGameEnd, Winner, Params, SystemInfo, Game, NumberCommit, NumberReveal };

async function initTxClient(vuexGetters) {
	return await txClient(vuexGetters['common/wallet/signer'], {
		addr: vuexGetters['common/env/apiTendermint']
	})
}

async function initQueryClient(vuexGetters) {
	return await queryClient({
		addr: vuexGetters['common/env/apiCosmos']
	})
}

function mergeResults(value, next_values) {
	for (let prop of Object.keys(next_values)) {
		if (Array.isArray(next_values[prop])) {
			value[prop]=[...value[prop], ...next_values[prop]]
		}else{
			value[prop]=next_values[prop]
		}
	}
	return value
}

function getStructure(template) {
	let structure = { fields: [] }
	for (const [key, value] of Object.entries(template)) {
		let field: any = {}
		field.name = key
		field.type = typeof value
		structure.fields.push(field)
	}
	return structure
}

const getDefaultState = () => {
	return {
				GetGame: {},
				GetGames: {},
				GetSystemInfo: {},
				GetParams: {},
				
				_Structure: {
						EventNewGame: getStructure(EventNewGame.fromPartial({})),
						EventRevealTimeout: getStructure(EventRevealTimeout.fromPartial({})),
						EventGameEnd: getStructure(EventGameEnd.fromPartial({})),
						Winner: getStructure(Winner.fromPartial({})),
						Params: getStructure(Params.fromPartial({})),
						SystemInfo: getStructure(SystemInfo.fromPartial({})),
						Game: getStructure(Game.fromPartial({})),
						NumberCommit: getStructure(NumberCommit.fromPartial({})),
						NumberReveal: getStructure(NumberReveal.fromPartial({})),
						
		},
		_Registry: registry,
		_Subscriptions: new Set(),
	}
}

// initial state
const state = getDefaultState()

export default {
	namespaced: true,
	state,
	mutations: {
		RESET_STATE(state) {
			Object.assign(state, getDefaultState())
		},
		QUERY(state, { query, key, value }) {
			state[query][JSON.stringify(key)] = value
		},
		SUBSCRIBE(state, subscription) {
			state._Subscriptions.add(JSON.stringify(subscription))
		},
		UNSUBSCRIBE(state, subscription) {
			state._Subscriptions.delete(JSON.stringify(subscription))
		}
	},
	getters: {
				getGetGame: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.GetGame[JSON.stringify(params)] ?? {}
		},
				getGetGames: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.GetGames[JSON.stringify(params)] ?? {}
		},
				getGetSystemInfo: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.GetSystemInfo[JSON.stringify(params)] ?? {}
		},
				getGetParams: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.GetParams[JSON.stringify(params)] ?? {}
		},
				
		getTypeStructure: (state) => (type) => {
			return state._Structure[type].fields
		},
		getRegistry: (state) => {
			return state._Registry
		}
	},
	actions: {
		init({ dispatch, rootGetters }) {
			console.log('Vuex module: zale144.whichnumber.whichnumber initialized!')
			if (rootGetters['common/env/client']) {
				rootGetters['common/env/client'].on('newblock', () => {
					dispatch('StoreUpdate')
				})
			}
		},
		resetState({ commit }) {
			commit('RESET_STATE')
		},
		unsubscribe({ commit }, subscription) {
			commit('UNSUBSCRIBE', subscription)
		},
		async StoreUpdate({ state, dispatch }) {
			state._Subscriptions.forEach(async (subscription) => {
				try {
					const sub=JSON.parse(subscription)
					await dispatch(sub.action, sub.payload)
				}catch(e) {
					throw new Error('Subscriptions: ' + e.message)
				}
			})
		},
		
		
		
		 		
		
		
		async QueryGetGame({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryGetGame( key.id)).data
				
					
				commit('QUERY', { query: 'GetGame', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryGetGame', payload: { options: { all }, params: {...key},query }})
				return getters['getGetGame']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryGetGame API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryGetGames({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryGetGames(query)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await queryClient.queryGetGames({...query, 'pagination.key':(<any> value).pagination.next_key})).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'GetGames', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryGetGames', payload: { options: { all }, params: {...key},query }})
				return getters['getGetGames']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryGetGames API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryGetSystemInfo({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryGetSystemInfo()).data
				
					
				commit('QUERY', { query: 'GetSystemInfo', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryGetSystemInfo', payload: { options: { all }, params: {...key},query }})
				return getters['getGetSystemInfo']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryGetSystemInfo API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryGetParams({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryGetParams()).data
				
					
				commit('QUERY', { query: 'GetParams', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryGetParams', payload: { options: { all }, params: {...key},query }})
				return getters['getGetParams']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryGetParams API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		async sendMsgRevealNumber({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgRevealNumber(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgRevealNumber:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgRevealNumber:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgUpdateParams({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgUpdateParams(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgUpdateParams:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgUpdateParams:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgCommitNumber({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgCommitNumber(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCommitNumber:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgCommitNumber:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgNewGame({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgNewGame(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgNewGame:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgNewGame:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		
		async MsgRevealNumber({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgRevealNumber(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgRevealNumber:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgRevealNumber:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgUpdateParams({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgUpdateParams(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgUpdateParams:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgUpdateParams:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgCommitNumber({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgCommitNumber(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCommitNumber:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgCommitNumber:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgNewGame({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgNewGame(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgNewGame:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgNewGame:Create Could not create message: ' + e.message)
				}
			}
		},
		
	}
}
