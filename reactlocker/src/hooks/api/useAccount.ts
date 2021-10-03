import { Account } from "../../shared/types";
import { useAccountList } from "./useAccountList";

export const useAccount = (id: number, enabled = true): [boolean, Account | null] => {
	const [isLoadingAccounts, accounts] = useAccountList();

	let result: [boolean, Account | null] = [true, null];

	if (isLoadingAccounts || accounts === []) {
		return result;
	}

	accounts.forEach((account) => {
		if (account.id === id) {
			result = [false, account];
		}
	});

	return result;
}