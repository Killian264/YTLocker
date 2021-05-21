import React from "react";
import { Card } from "./Card";
import { LoadingStatsCard } from "./LoadingStatsCard";

export interface LoadingUserInfoBarProps {
	className?: string;
}

export const LoadingUserInfoBar: React.FC<LoadingUserInfoBarProps> = ({ className, ...props }) => {
	return (
		<div className={`${className} flex`} {...props}>
			<div className="bg-primary-700 p-32 rounded-md mr-3 hidden md:block"></div>
			<div className="flex-grow flex items-end">
				<Card className="flex-grow">
					<BarHeader />
					<div className="grid xs:grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-2.5 max-w-7xl">
						<LoadingStatsCard />
						<LoadingStatsCard />
						<LoadingStatsCard />
						<LoadingStatsCard />
					</div>
				</Card>
			</div>
		</div>
	);
};

const BarHeader: React.FC<{}> = () => {
	const color = "bg-primary-600 text-primary-600 rounded leading-5";

	return (
		<div className="flex justify-between">
			<div>
				<span
					className={`${color} text-2xl font-semibold inline-block align-bottom leading-none animate-pulse`}
				>
					asdfasdfasdfasdf
				</span>
			</div>
			<div className="flex flex-col justify-end animate-pulse">
				<span className={`${color} text-sm leading-none`}>Joined March 17th 2020</span>
			</div>
		</div>
	);
};
