import React from "react";
import { StatsCard } from "./StatsCard";
import { Badge } from "./Badge";
import { Card } from "./Card";
import { StatCard, User } from "../shared/types";

export interface UserInfoBarProps {
	className?: string;
	user: User;
	stats: StatCard[];
}

export const UserInfoBar: React.FC<UserInfoBarProps> = ({
	user,
	stats,
	className,
	...props
}) => {
	return (
		<div className={`${className} flex`} {...props}>
			<UserProfile />
			<div className="flex-grow flex items-end">
				<Card className="flex-grow">
					<BarHeader user={user} />
					<MultiCard stats={stats} />
				</Card>
			</div>
		</div>
	);
};

interface BarHeaderProps {
	user: User;
}

const BarHeader: React.FC<BarHeaderProps> = ({ user }) => {
	return (
		<div className="flex justify-between">
			<div>
				<span className=" text-2xl font-semibold inline-block align-bottom leading-none">
					{user.username}
				</span>
				<Badge className="ml-1 mt-3" color="primary">
					PRO
				</Badge>
			</div>
			<div className="flex flex-col justify-end">
				<span className="text-sm leading-none">{`Joined ${user.joined}`}</span>
			</div>
		</div>
	);
};

interface MultiCardProps {
	stats: StatCard[];
}

const MultiCard: React.FC<MultiCardProps> = ({ stats }) => {
	const cards = stats.map((stat, index) => {
		return <StatsCard key={index} {...stat} />;
	});

	return (
		<div className="grid xs:grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-2.5 max-w-7xl">
			{cards}
		</div>
	);
};

const UserProfile: React.FC<{}> = () => {
	return (
		<div className="bg-primary-700 p-32 rounded-md mr-3 hidden md:block">
			{" "}
		</div>
	);
};
