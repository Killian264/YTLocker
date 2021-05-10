import React from "react";
import PropTypes from "prop-types";
import { StatsCard } from "./StatsCard";
import { Badge } from "./Badge";
import { Card } from "./Card";
import { StatCard, User } from "../shared/types";

interface UserInfoBarProps {
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

interface BarHeader {
	user: User;
}

const BarHeader: React.FC<BarHeader> = ({ user }) => {
	return (
		<div className="flex justify-between">
			<div>
				<span className=" text-2xl inline-block align-bottom leading-none">
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

interface MultiCard {
	stats: StatCard[];
}

const MultiCard: React.FC<MultiCard> = ({ stats }) => {
	const cards = stats.map((stat, index) => {
		return <StatsCard key={index} {...stat} />;
	});

	return (
		<div className="grid xs:grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 max-w-7xl">
			{cards}
		</div>
	);
};

const UserProfile: React.FC<{}> = () => {
	return (
		<div className="bg-secondary p-32 rounded-md mr-3 hidden md:block">
			{" "}
		</div>
	);
};
