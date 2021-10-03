import { useState } from "react";
import { Account, Playlist } from "../shared/types";
import { BuildPlaylistUrl } from "../shared/urls";
import { Button } from "./Button";
import { Card } from "./Card";
import { ColorBadge } from "./ColorBadge";
import { Dropdown } from "./Dropdown";
import { Modal } from "./Modal";
import { Cancel, Cog, Copy, ExternalLink, Pause, Reload, SvgBox, Trash } from "./Svg";
import { TextBox } from "./TextBox";

export interface PlaylistViewProps {
	className?: string;
	playlist: Playlist;
	account: Account;
	accounts: Account[];
	EditClick: () => void;
	DeleteClick: () => void;
	PauseClick: () => void;
	CopyClick: () => void;
	RefreshClick: () => void;
	BackClick: () => void;
}

export const PlaylistView: React.FC<PlaylistViewProps> = ({
	className,
	playlist,
	account,
	accounts,
	EditClick,
	DeleteClick,
	PauseClick,
	CopyClick,
	RefreshClick,
	BackClick,
}) => {
	const [isOpenDelete, setIsOpenDelete] = useState(false);
	const [isOpenPause, setIsOpenPause] = useState(false);
	const [isOpenCopy, setIsOpenCopy] = useState(false);
	const [isOpenRefresh, setIsOpenRefresh] = useState(false);

	let deleteModal = (
		<Modal
			confirmMessage={"Yes, I am sure"}
			header={"Are you sure?"}
			body={"Playlist information will be deleted, but the playlist will remain accessible on youtube."}
			AcceptClick={() => {
				DeleteClick();
				BackClick();
			}}
			RejectClick={() => {
				setIsOpenDelete(false);
			}}
		/>
	);

	let pauseModal = (
		<Modal
			header={"Are you sure?"}
			body={"Pausing the playlist will stop new videos from being added."}
			AcceptClick={() => {
				PauseClick();
				setIsOpenPause(false);
			}}
			RejectClick={() => {
				setIsOpenPause(false);
			}}
		/>
	);

	let copyModal = (
		<Modal
			header={"Are you sure?"}
			body={
				"Copying will create a new playlist with the same information and subscriptions but not videos."
			}
			AcceptClick={() => {
				CopyClick();
				BackClick();
			}}
			RejectClick={() => {
				setIsOpenCopy(false);
			}}
		/>
	);

	let refreshModal = (
		<Modal
			header={"Are you sure?"}
			body={
				"Refreshing will refresh the playlist to match what is on Youtube. Only needs to be ran after a user manually updates a playlist."
			}
			AcceptClick={() => {
				RefreshClick();
				setIsOpenRefresh(false);
			}}
			RejectClick={() => {
				setIsOpenCopy(false);
			}}
		/>
	);

	return (
		<div>
			{isOpenDelete ? deleteModal : <div></div>}
			{isOpenPause ? pauseModal : <div></div>}
			{isOpenCopy ? copyModal : <div></div>}
			{isOpenRefresh ? refreshModal : <div></div>}
			<Card className={`${className} flex flex-col`}>
				<div className="flex justify-between -mb-2 -mt-2">
					<div className="text-2xl font-semibold mt-auto">
						<span className="leading-none -mt-0.5">View Playlist</span>
					</div>
					<div className="flex gap-2">
						<SvgBox className={`border-primary-200 p-0.5`} onClick={EditClick}>
							<Cog className="text-primary-200" size={24}></Cog>
						</SvgBox>
					</div>
				</div>
				<a
					className="relative"
					href={BuildPlaylistUrl(playlist.youtubeId)}
					target="_blank"
					rel="noreferrer"
				>
					<div className="absolute right-0 opacity-20">
						<ExternalLink size={38}></ExternalLink>
					</div>
					<div className={`col-span-6 rounded-lg object-cover w-full bg-black h-40`} />
				</a>
				<div className="flex gap-2 mt-3">
					<div className="flex-grow">
						<div className="flex flex-row flex gap-2 justify-between">
							<div className="flex">
								<span className="md:text-2xl text-2xl font-semibold block my-auto">
									{playlist.title}
								</span>
								<div className="my-auto ml-2">
									{playlist.isActive ? (
										<SvgBox className={`flex text-green-300 border-green-300`}>
											<span className="text-md font-semibold mx-1">Active</span>
										</SvgBox>
									) : (
										<SvgBox className={`flex text-black border-black`}>
											<span className="text-md font-semibold mx-1">Paused</span>
										</SvgBox>
									)}
								</div>
								<div className="m-auto ml-3 mt-3">
									<ColorBadge color={playlist.color}></ColorBadge>
								</div>
							</div>
							<Dropdown
								className="mt-3 mb-2 sm:m-0"
								defaultSelected={account.username}
								title="Account Selection"
								items={[]}
								disabled={true}
							></Dropdown>
						</div>
						<TextBox
							className="my-3"
							disabled={true}
							placeholder="Description"
							value={playlist.description}
						></TextBox>
					</div>
				</div>
				<div className="flex justify-between mt-auto pt-1">
					<Button size="medium" color="secondary" onClick={BackClick}>
						Back
					</Button>
					<div className="flex gap-2 my-auto">
						<div className="gap-2 flex">
							<SvgBox
								className={`p-0.5 flex`}
								onClick={() => {
									setIsOpenRefresh(true);
								}}
							>
								<Reload size={24}></Reload>
								<span className="font-semibold mx-1">Refresh</span>
							</SvgBox>
						</div>
						<div className="gap-2 flex">
							<SvgBox
								className={`p-0.5 flex`}
								onClick={() => {
									setIsOpenCopy(true);
								}}
							>
								<Copy size={24}></Copy>
								<span className="font-semibold mx-1">Copy</span>
							</SvgBox>
						</div>
						<div className="gap-2 flex">
							<SvgBox
								className={`p-0.5 flex`}
								onClick={() => {
									setIsOpenPause(true);
								}}
							>
								<Pause size={24}></Pause>
								<span className="font-semibold mr-1">Pause</span>
							</SvgBox>
						</div>
						<div className="gap-2 flex">
							<SvgBox
								className={`p-0.5 flex text-red-500 border-red-500`}
								onClick={() => {
									setIsOpenDelete(true);
								}}
							>
								<Trash size={24}></Trash>
								<span className="text-red-500 font-semibold mx-1">Delete</span>
							</SvgBox>
						</div>
					</div>
				</div>
			</Card>
		</div>
	);
};
