import React from "react";
import { StatCard } from "../shared/types";

export interface StatsCardProps extends StatCard {
	className?: string;
}

const border = "border-t-8 border-accent border-solid";

export const StatsCard: React.FC<StatsCardProps> = ({
	className = "",
	header,
	count,
	measurement,
	...props
}) => {
	return (
		<div className={`bg-primary-600 px-4 pb-2 pt-2 rounded-md ${border} ${className}`} {...props}>
			<span className="text-lg font-semibold">{header}</span>
			<div>
				<span className="text-xl text-accent">{count}</span>
				<span className="text-lg"> {measurement}</span>
			</div>
		</div>
	);
};
