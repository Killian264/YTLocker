import React from "react";

export interface LoadingStatsCardProps {
	className?: string;
}

const border = "border-t-8 border-accent border-solid";

export const LoadingStatsCard: React.FC<LoadingStatsCardProps> = ({ className = "" }) => {
	const color = "bg-primary-700 text-primary-700 rounded leading-5";

	return (
		<div className={`bg-primary-600 px-4 pb-2 pt-2 rounded-md ${border} ${className}`}>
			<span className={`${color} text-lg font-semibold animate-pulse`}>Subscriptionsasdfasdf</span>
			<div className="mt-2 animate-pulse">
				<span className={`${color} text-xl mr-1`}>42</span>
				<span className={`${color} text-lg`}>secondsaa</span>
			</div>
		</div>
	);
};
