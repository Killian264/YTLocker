import React from "react";
import { StatsCard } from "./StatsCard";
import { Badge } from "./Badge";
import { Card } from "./Card";
import { StatCard, User } from "../shared/types";
import { Logo } from "./Logo";
import { Link } from "./Link";

export interface UserInfoBarProps {
	className?: string;
	HomeClick: () => void;
	LogOutClick: () => void;
	user: User;
	stats: StatCard[];
}

export const UserInfoBar: React.FC<UserInfoBarProps> = ({
	user,
	stats,
	className,
	HomeClick,
	LogOutClick,
	...props
}) => {
	return (
		<div className={`${className} flex`} {...props}>
			<UserProfile />
			<div className="flex flex-col flex-grow gap-2 justify-end">
				<div className="flex justify-between gap-4">
					<div className="flex gap-2">
						<Logo></Logo>
						<span className="text-xl font-semibold">YTLocker</span>
					</div>
					<div className="gap-6 flex mt-auto">
						<Link className="text-lg font-semibold" onClick={HomeClick}>
							Home
						</Link>
						<Link className="sm:block hidden text-lg font-semibold underline">Dashboard</Link>
						<Link className="text-lg font-semibold" onClick={LogOutClick}>
							Log Out
						</Link>
					</div>
				</div>
				<div className="flex">
					<Card className="flex-grow">
						<BarHeader user={user} />
						<MultiCard stats={stats} />
					</Card>
				</div>
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
				<span className="text-sm leading-none">{`Joined ${user.joined.toDateString()}`}</span>
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

	return <div className="grid xs:grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-2.5 max-w-7xl">{cards}</div>;
};

const UserProfile: React.FC<{}> = () => {
	return <div className="bg-primary-700 p-32 rounded-md mr-3 hidden md:block"> </div>;
};
