import { Card } from "../components/Card";
import { GoogleOAuthCard } from "../components/GoogleOauthCard";
import { BeforeLinkingInfo, BaseAccountInfo } from "../components/InfoCards";
import { LoadingListItem } from "../components/LoadingListItem";
import { AccountListItemController } from "../controllers/AccountListController";
import { useAccountList } from "../hooks/api/useAccountList";

export const DashboardAccountsView: React.FC<{}> = () => {
	const [isLoadingAccounts, accounts] = useAccountList();

	let items = [
		<LoadingListItem></LoadingListItem>,
		<LoadingListItem></LoadingListItem>
	]

	if (!isLoadingAccounts) {
		items = accounts.map((accountId) => {
			return <AccountListItemController
				accountId={accountId}
			></AccountListItemController>
		})
	}

	return (
		<div className="w-100">
			<div className="flex w-100 gap-8 flex-col lg:flex-row">
				<div className="flex-grow flex flex-col gap-8 lg:w-2/3">
					<Card>
						<div className="flex justify-between -mb-2 -mt-2">
							<div className="text-2xl font-semibold">
								<span className="leading-none -mt-0.5">Linked Accounts</span>
							</div>
						</div>
						<div>
							{items}
						</div>
					</Card>
					<BeforeLinkingInfo></BeforeLinkingInfo>
				</div>
				<div className="lg:w-1/3">
					<div className="flex flex-col md:flex-row lg:flex-col gap-8 items-start">
						<GoogleOAuthCard className="mx-auto flex-grow" type="link"></GoogleOAuthCard>
						<BaseAccountInfo></BaseAccountInfo>
					</div>
				</div>
			</div>
		</div>
	);
};