import { Account } from "../shared/types";
import { Badge } from "./Badge";
import { RightArrow } from "./Svg";

export interface AccountListItemProps {
	className?: string;
	account: Account;
}

export const AccountListItem: React.FC<AccountListItemProps> = ({className, account}) => {
	const css = `hover:bg-primary-600 rounded-md flex justify-between cursor-pointer overflow-hidden`;
	const imageSize = "md:h-20 md:w-32 h-16 w-24";
	const textSize = "sm:text-md text-md";

	return (
		<div className={className + css} onClick={() => {}}>
		<div className="flex p-1 overflow-hidden">
			<div
				className={`rounded-lg flex-shrink-0 bg-black ${imageSize} flex justify-center items-center`}
			>
				<img src={account.picture} className="w-12 rounded"></img>
			</div>
			<div className="pl-3 flex flex-col">
				<span className={`${textSize} font-semibold whitespace-nowrap`}>
					{account.username}
					{/* <span className="text-accent ml-2">{3}</span>/6 */}
				</span>
				<span className="whitespace-nowrap whitespace-nowrap mt-0.5">
					{/* <Badge className="mt-2 mr-2" color="primary">
						YTLocker
					</Badge> */}
					<Badge className="mr-2" color="secondary">
						{account.permissionLevel == "view" ? "VIEW-ONLY" : "CREATION"}
					</Badge>
					{/* <Badge color="secondary">
						LIMITED
					</Badge> */}
					<div className="mt-0.5">
						<span className="font-semibold">Email: </span> {account.email}
					</div>
				</span>
			</div>
		</div>
		<div className="mr-2 my-auto select-none">
			<RightArrow size={24}></RightArrow>
		</div>
	</div>
	)
}