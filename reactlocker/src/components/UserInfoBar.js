import React from 'react';
import PropTypes from 'prop-types';
import { StatsCard } from './StatsCard';
import { Badge } from './Badge';
import { Card } from './Card';

const UserShape = PropTypes.shape({
	username: PropTypes.string,
	email: PropTypes.string,
	joined: PropTypes.string,
});

const StatsCardShape = PropTypes.shape({
	header: PropTypes.string,
	count: PropTypes.number,
	measurement: PropTypes.string,
})

export const UserInfoBar = ({ user, statCards, className, ...props }) => {
	return (
		<div className={`${className} flex`} {...props}>
			<UserProfile/>
			<div className="max-w-7xl flex-grow flex items-end">
				<Card className="flex-grow" >
					<BarHeader user={user} />
					<MultiCard cards={statCards} />
				</Card>
			</div>
		</div>
	);
};
UserInfoBar.propTypes = {
	className: PropTypes.string,
	user: UserShape,
	statCards: PropTypes.arrayOf(StatsCardShape)
};

const BarHeader = ( {user} ) => {
	return (
		<div className="flex justify-between">
			<div>
				<span className=" text-2xl inline-block align-bottom leading-none">{user.username}</span>
				<Badge className="ml-1 mt-3" >PRO</Badge>
			</div>
			<div className="flex flex-col justify-end">
				<span className="text-sm leading-none">{`Joined ${user.joined}`}</span>
			</div>
		</div>
	)
}
BarHeader.propTypes = {
	user: UserShape,
}

const MultiCard = ({ cards }) => {
	const stats = [];

	cards.forEach( (card, index) => {
		stats.push(<StatsCard key={index} {...card} />)
	} )

	return (
		<div className="grid xs:grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 max-w-7xl">
			{stats}
		</div>
	)
}
MultiCard.propTypes = {
	cards: PropTypes.arrayOf(StatsCardShape)
}

const UserProfile = () => {
	return (<div className="bg-secondary p-32 rounded-md mr-3 hidden md:block" > </div>)
}