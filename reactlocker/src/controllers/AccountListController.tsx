import React from "react";
import { AccountListItem } from "../components/AccountListItem";
import { LoadingListItem } from "../components/LoadingListItem";
import { useAccount } from "../hooks/api/useAccount";

export interface AccountListItemControllerProps {
	className?: string;
	accountId: number;
}

export const AccountListItemController: React.FC<AccountListItemControllerProps> = ({
	className,
	accountId,
}) => {
	const [isLoadingAccount, account] = useAccount(accountId);

	if (isLoadingAccount || account == null) {
		return <LoadingListItem></LoadingListItem>;
	}

	return (
		<AccountListItem
			className={className}
			account={account}
		></AccountListItem>
	);
};
