import { useState } from "react";
import { Button } from "../components/Button";
import { Card } from "../components/Card";
import { Input } from "../components/Input";
import { PlaylistColorSelector } from "./PlaylistColorSelector";
import { Account, Color, Playlist } from "../shared/types";
import { TextBox } from "./TextBox";
import { Dropdown } from "./Dropdown";

export interface ChannelSubscribeCardProps {
	className?: string;
	editPlaylist?: Playlist | null;
	playlists: Playlist[];
	accounts: Account[];
	CreateClick: (title: string, description: string, color: Color, accountId: number) => void;
	BackClick: () => void;
}

export const PlaylistCreateCard: React.FC<ChannelSubscribeCardProps> = ({
	className,
	editPlaylist = null,
	playlists,
	accounts,
	CreateClick,
	BackClick,
}) => {
	const [title, setTitle] = useState(editPlaylist === null ? "" : editPlaylist.title);
	const [description, setDescription] = useState(editPlaylist === null ? "" : editPlaylist.description);
	const [color, setColor] = useState<Color | null>(editPlaylist === null ? null : editPlaylist.color);
	const [accountId, setAccountId] = useState(editPlaylist === null ? null : editPlaylist.accountId);

	let defaultSelected = "";

	const accountDropdownItems = accounts.map((account) => {
		if (editPlaylist !== null && account.id == editPlaylist.accountId) {
			defaultSelected = account.username;
		}

		return {
			title: account.username,
			value: account.id,
		};
	});

	return (
		<Card className={`${className} flex flex-col justify-content-between`}>
			<div className="flex justify-between -mb-1 -mt-1">
				<span className="leading-none text-2xl font-semibold">
					{editPlaylist === null ? "Create" : "Edit"} Playlist
				</span>
			</div>
			<div className={`col-span-6 rounded-lg object-cover w-full bg-black h-40`} />
			<div className="mt-3">
				<div className="flex flex-col sm:flex-row mb-3 justify-between">
					<div className="flex">
						<span className="text-xl font-semibold my-auto pb-1">Color:</span>
						<div className="my-auto ml-3">
							<PlaylistColorSelector
								selectedPlaylist={editPlaylist}
								OnClick={(color) => {
									setColor(color);
								}}
								playlists={playlists}
							></PlaylistColorSelector>
						</div>
					</div>
					<Dropdown
						className="mt-3 mb-2 sm:m-0"
						OnItemSelected={(accountId) => {
							setAccountId(accountId);
						}}
						defaultSelected={defaultSelected}
						title="Account Selection"
						items={accountDropdownItems}
						disabled={defaultSelected !== ""}
					></Dropdown>
				</div>
				<Input
					placeholder="Title"
					value={title}
					className="mb-3"
					onChange={(e: React.ChangeEvent<HTMLInputElement>) => {
						setTitle(e.target.value);
					}}
				></Input>
				<TextBox
					disabled={true}
					placeholder="Description"
					value={description}
					className="mb-3"
					onChange={(e: React.ChangeEvent<HTMLTextAreaElement>) => {
						setDescription(e.target.value);
					}}
				></TextBox>
				<div>
					<Button
						size="medium"
						color="primary"
						disabled={title === "" || color === null || accountId === null}
						onClick={() => {
							if (title === "" || color === null || accountId === null) {
								return;
							}

							CreateClick(title, description, color, accountId);
							BackClick();
						}}
					>
						{editPlaylist === null ? "Create" : "Save"}
					</Button>
				</div>
				<div className="pt-4 flex justify-between">
					<Button size="medium" color="secondary" onClick={BackClick}>
						Back
					</Button>
					<Button
						size="medium"
						color="primary"
						disabled={title === "" || color === null || accountId === null}
						onClick={() => {
							if (title === "" || color === null || accountId === null) {
								return;
							}

							CreateClick(title, description, color, accountId);
							BackClick();
						}}
					>
						{editPlaylist === null ? "Create" : "Save"}
					</Button>
				</div>
			</div>
		</Card>
	);
};
