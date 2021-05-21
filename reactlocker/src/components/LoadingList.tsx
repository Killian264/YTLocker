import React from "react";
import { Card } from "./Card";
import { LoadingListItem } from "./LoadingListItem";

export interface LoadingListProps {
	className?: string;
	limit: number;
}

export const LoadingList: React.FC<LoadingListProps> = ({ className, limit }) => {
	const color = "bg-primary-600 text-primary-600 rounded leading-5";

	let list: JSX.Element[] = [];
	for (let i = 0; i < limit; i++) {
		list.push(<LoadingListItem key={i}></LoadingListItem>);
	}

	return (
		<Card className={className}>
			<div className="flex justify-between -mb-1 -mt-1">
				<div className="text-2xl font-semibold">
					<span className={`${color} leading-none -mt-0.5`}>asdfasdkfjasdkfjasf</span>
				</div>
			</div>
			<div className="grid gap-2">{list}</div>
		</Card>
	);
};
