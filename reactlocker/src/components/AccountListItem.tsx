import { Account } from "../shared/types";
import { Badge } from "./Badge";

export interface AccountListItemProps {
	className?: string;
	account: Account;
}

export const AccountListItem: React.FC<AccountListItemProps> = ({ className, account }) => {
	const css = `hover:bg-primary-600 rounded-md flex justify-between overflow-hidden`;
	const imageSize = "md:h-20 md:w-32 h-16 w-24";
	const textSize = "sm:text-md text-md";

	return (
		<div className={className + css} onClick={() => {}}>
			<div className="flex p-1 overflow-hidden">
				<div
					className={`rounded-lg flex-shrink-0 bg-black ${imageSize} flex justify-center items-center`}
				>
					<img src={account.picture} className="w-12 rounded" alt="youtube account" />
				</div>
				<div className="pl-3 flex flex-col">
					<span className={`${textSize} font-semibold whitespace-nowrap`}>{account.username}</span>
					<span className="whitespace-nowrap whitespace-nowrap mt-0.5">
						<Badge className="mr-2" color="secondary">
							{account.permissionLevel === "view" ? "VIEW-ONLY" : "CREATION"}
						</Badge>
						<div className="mt-0.5">
							<span className="font-semibold">Email: </span> {account.email}
						</div>
					</span>
				</div>
			</div>
			<div className="mr-2 ml-4 my-auto select-none"></div>
		</div>
	);
};
